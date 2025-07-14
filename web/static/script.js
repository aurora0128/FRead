// 全局变量
let articles = [];
let categories = ['技术', '生活', '美食'];
let currentArticle = null;
let currentCategory = 'all';

// API 配置
const API_BASE_URL = 'http://localhost:8080/api'; // 请根据您的后端地址修改

// 初始化
document.addEventListener('DOMContentLoaded', function() {
    loadArticles();
    loadCategories();
    setupEventListeners();
});

// 设置事件监听器
function setupEventListeners() {
    // 分类筛选
    document.getElementById('categoryList').addEventListener('click', function(e) {
        if (e.target.classList.contains('category-tag')) {
            // 移除所有激活状态
            document.querySelectorAll('.category-tag').forEach(tag => {
                tag.classList.remove('active');
            });
            // 添加激活状态
            e.target.classList.add('active');
            currentCategory = e.target.dataset.category;
            filterArticles();
        }
    });

    // 回车键添加文章
    document.getElementById('urlInput').addEventListener('keypress', function(e) {
        if (e.key === 'Enter') {
            addArticle();
        }
    });

    // 回车键添加分类
    document.getElementById('categoryInput').addEventListener('keypress', function(e) {
        if (e.key === 'Enter') {
            addCategory();
        }
    });
}

// 添加文章
async function addArticle() {
    const url = document.getElementById('urlInput').value.trim();
    if (!url) {
        alert('请输入链接');
        return;
    }
    try {
        showLoading();
        const response = await fetch(`${API_BASE_URL}/articles`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ url: url })
        });

        if (!response.ok) {
            throw new Error('添加失败');
        }

        const article = await response.json();
        articles.unshift(article.data);
        // loadArticles()
        renderArticles()
        document.getElementById('urlInput').value = '';
        hideLoading();
    } catch (error) {
        console.error('添加文章失败:', error);
        alert('添加文章失败，请重试');
        hideLoading();
    }
}

// 加载文章列表
async function loadArticles() {
    try {
        const response = await fetch(`${API_BASE_URL}/articles`);
        if (response.ok) {
            const responseData = await response.json();

            if (responseData && Array.isArray(responseData.data)) {
                articles = responseData.data;
                renderArticles();
            } else {
                console.error('服务器返回的数据格式不正确，未找到文章数组。', responseData);
                articles = []; // 确保 articles 仍然是数组，避免后续报错
            }
        }
    } catch (error) {
        console.error('加载文章失败:', error);
    }
}

// 加载分类
async function loadCategories() {
    try {
        const response = await fetch(`${API_BASE_URL}/categories`);
        if (response.ok) {
            categories = await response.json();
            renderCategories();
        }
    } catch (error) {
        console.error('加载分类失败:', error);
    }
}

// 渲染文章列表
// function renderArticles() {
//     const container = document.getElementById('articlesList');
//     const filteredArticles = currentCategory === 'all' ?
//         articles : articles.filter(article => article.data.category === currentCategory);
//
//     container.innerHTML = filteredArticles.map(article => `
//             <div class="article-item" onclick="selectArticle('${article.id}')">
//                 <div class="article-title">${article.title || '无标题'}</div>
//                 <div class="article-url">${article.url}</div>
//                 <div class="article-date">${formatDate(article.createdAt)}</div>
//             </div>
//         `).join('');
// }

function renderArticles() {
    const container = document.getElementById('articlesList');
    const filteredArticles = currentCategory === 'all' ?
        articles : articles.filter(article => article.category === currentCategory);

    container.innerHTML = filteredArticles.map(article => `
        <div class="article-item" onclick="selectArticle('${article.id}')">
            <div class="article-content">
                <div class="article-info">
                    <div class="article-title">${article.title || '无标题'}</div>
                    <div class="article-date">${formatDate(article.createdAt)}</div>
                </div>
                <div class="article-thumbnail">
<!--                todo:添加默认缩略图-->
                    <img src="${article.thumbnail || '/default-thumbnail.jpg'}" 
                         alt="${article.title || '无标题'}" 
                         onerror="this.src='/default-thumbnail.jpg'">
                </div>
            </div>
        </div>
    `).join('');
}


// 渲染分类
function renderCategories() {
    const container = document.getElementById('categoryList');
    const allTag = '<div class="category-tag active" data-category="all">全部</div>';
    const categoryTags = categories.map(category =>
        `<div class="category-tag" data-category="${category}">${category}</div>`
    ).join('');
    container.innerHTML = allTag + categoryTags;
}

// 筛选文章
function filterArticles() {
    renderArticles();
}

// 选择文章
async function selectArticle(articleId) {
    try {
        // 更新界面状态
        document.querySelectorAll('.article-item').forEach(item => {
            item.classList.remove('active');
        });
        event.target.closest('.article-item').classList.add('active');

        // 显示加载状态
        showContentLoading();

        // 获取文章详情
        const response = await fetch(`${API_BASE_URL}/articles/${articleId}`);
        if (!response.ok) {
            throw new Error('获取文章失败');
        }

        const article = await response.json();
        currentArticle = article;

        // 更新内容区域
        document.getElementById('contentTitle').textContent = article.data.title || '无标题';

        // 使用marked库将Markdown转换为HTML
        const markdownContent = article.data.content || '暂无内容';
        const htmlContent = marked.parse(markdownContent);

        document.getElementById('contentBody').innerHTML = `
            <div class="markdown-content">
                ${htmlContent}
            </div>
        `;
    } catch (error) {
        console.error('选择文章失败:', error);
        showError('加载文章失败，请重试');
    }
}
// 添加分类
async function addCategory() {
    const categoryName = document.getElementById('categoryInput').value.trim();
    if (!categoryName) {
        alert('请输入分类名称');
        return;
    }

    if (categories.includes(categoryName)) {
        alert('分类已存在');
        return;
    }

    try {
        const response = await fetch(`${API_BASE_URL}/categories`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name: categoryName })
        });

        if (!response.ok) {
            throw new Error('添加分类失败');
        }

        categories.push(categoryName);
        renderCategories();
        document.getElementById('categoryInput').value = '';
    } catch (error) {
        console.error('添加分类失败:', error);
        alert('添加分类失败，请重试');
    }
}

// 删除文章
async function deleteArticle() {
    if (!currentArticle) {
        alert('请选择要删除的文章');
        return;
    }

    if (!confirm('确定要删除这篇文章吗？')) {
        return;
    }

    try {
        const response = await fetch(`${API_BASE_URL}/articles/${currentArticle.data.id}`, {
            method: 'DELETE'
        });

        if (!response.ok) {
            throw new Error('删除失败');
        }

        articles = articles.filter(article => article.id !== currentArticle.data.id);
        renderArticles();

        // 清空内容区域
        document.getElementById('contentTitle').textContent = '选择文章查看内容';
        document.getElementById('contentBody').innerHTML = `
                <div class="empty-state">
                    <div class="empty-icon">📄</div>
                    <div>请从左侧选择文章查看内容</div>
                </div>
            `;
        currentArticle = null;
    } catch (error) {
        console.error('删除文章失败:', error);
        alert('删除文章失败，请重试');
    }
}

// 工具函数
function formatDate(dateString) {
    const date = new Date(dateString);
    return date.toLocaleDateString('zh-CN', {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit'
    });
}

function showLoading() {
    // 这里可以显示全局加载状态
}

function hideLoading() {
    // 这里可以隐藏全局加载状态
}

function showContentLoading() {
    document.getElementById('contentBody').innerHTML = `
            <div class="loading">
                <div class="spinner"></div>
            </div>
        `;
}

function showError(message) {
    document.getElementById('contentBody').innerHTML = `
            <div class="empty-state">
                <div class="empty-icon">❌</div>
                <div>${message}</div>
            </div>
        `;
}