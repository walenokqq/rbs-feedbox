import { navigateTo } from '../utils/router';
import { getProjectForms } from '../utils/api';

export function renderProjectFormsPage(projectId) {
  const app = document.getElementById('app');
  
  app.innerHTML = `
    <div class="container">
      <h1>Формы проекта</h1>
      <div class="actions-bar">
        <button id="backToProjectsBtn" class="btn btn-secondary">← Назад к проектам</button>
        <button id="createFormBtn" class="btn btn-primary">+ Создать новую форму</button>
      </div>
      
      <table class="forms-table">
        <thead>
          <tr>
            <th>Название формы</th>
            <th>Количество ответов</th>
            <th>Дата последнего ответа</th>
            <th>Дата создания</th>
            <th>Действия</th>
          </tr>
        </thead>
        <tbody id="formsTableBody">
          <tr>
            <td colspan="5" class="loading-message">Загрузка данных...</td>
          </tr>
        </tbody>
      </table>
    </div>
  `;

  // Навигация назад
  document.getElementById('backToProjectsBtn').addEventListener('click', () => {
    navigateTo('/projects');
  });

  // Переход в конструктор форм
  document.getElementById('createFormBtn').addEventListener('click', () => {
    navigateTo(`/projects/${projectId}/forms/new`);
  });

  // Загружаем формы проекта
  loadProjectForms(projectId);
}

async function loadProjectForms(projectId) {
  try {
    const forms = await getProjectForms(projectId);
    const tableBody = document.getElementById('formsTableBody');
    
    if (forms.length === 0) {
      tableBody.innerHTML = `
        <tr>
          <td colspan="5" class="empty-message">В этом проекте пока нет форм</td>
        </tr>
      `;
      return;
    }
    
    tableBody.innerHTML = forms.map(form => `
      <tr>
        <td>${form.title || 'Без названия'}</td>
        <td>${form.responses_count || 0}</td>
        <td>${formatDateTime(form.last_response_date) || 'Нет ответов'}</td>
        <td>${formatDateTime(form.created_at)}</td>
        <td class="actions-cell">
          <button class="btn btn-primary view-responses-btn" data-id="${form.id}">
            Просмотр ответов
          </button>
        </td>
      </tr>
    `).join('');
    
    // Обработчики для кнопок просмотра ответов
    document.querySelectorAll('.view-responses-btn').forEach(btn => {
      btn.addEventListener('click', (e) => {
        const formId = e.target.getAttribute('data-id');
        navigateTo(`/feedback/${formId}`);
      });
    });
    
  } catch (error) {
    console.error('Ошибка при загрузке форм:', error);
    document.getElementById('formsTableBody').innerHTML = `
      <tr>
        <td colspan="5" class="error-message">Не удалось загрузить формы. Пожалуйста, попробуйте позже.</td>
      </tr>
    `;
  }
}

// Функция для форматирования даты и времени
function formatDateTime(dateString) {
  if (!dateString) return '';
  
  try {
    const date = new Date(dateString);
    return date.toLocaleString('ru-RU', {
      day: '2-digit',
      month: '2-digit',
      year: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });
  } catch (e) {
    console.error('Ошибка форматирования даты:', e);
    return dateString;
  }
}