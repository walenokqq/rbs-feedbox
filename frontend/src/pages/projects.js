import { navigateTo } from '../utils/router';
import { getProjects, createProject } from '../utils/api';
import { formatDate } from '../utils/helpers';

export function renderProjectsPage() {
  const app = document.getElementById('app');
  
  app.innerHTML = `
    <div class="container">
      <h1>Проекты 111</h1>
      <button id="createProjectBtn" class="btn btn-primary">Создать проект</button>
      
      <div class="project-form" style="display: none; margin: 20px 0;">
        <input type="text" id="projectTitle" placeholder="Название проекта" class="form-input">
        <button id="submitProjectBtn" class="btn btn-success">Подтвердить</button>
        <button id="cancelProjectBtn" class="btn btn-secondary">Отменить</button>
      </div>
      
      <table class="projects-table">
        <thead>
          <tr>
            <th>Название проекта</th>
            <th>Количество форм</th>
            <th>Последняя форма создана</th>
            <th>Проект создан</th>
            <th>Просмотр форм обратной связи</th>
          </tr>
        </thead>
        <tbody id="projectsTableBody"></tbody>
      </table>
    </div>
  `;

  const createProjectBtn = document.getElementById('createProjectBtn');
  const projectForm = document.querySelector('.project-form');
  const submitProjectBtn = document.getElementById('submitProjectBtn');
  const cancelProjectBtn = document.getElementById('cancelProjectBtn');
  const projectTitleInput = document.getElementById('projectTitle');
  
  createProjectBtn.addEventListener('click', () => {
    projectForm.style.display = 'block';
  });
  
  cancelProjectBtn.addEventListener('click', () => {
    projectForm.style.display = 'none';
    projectTitleInput.value = '';
  });
  
  submitProjectBtn.addEventListener('click', async () => {
    const title = projectTitleInput.value.trim();
    if (title) {
      await createProject(title);
      projectForm.style.display = 'none';
      projectTitleInput.value = '';
      loadProjects();
    }
  });
  
  loadProjects();
}

async function loadProjects() {
  const projects = await getProjects();
  const tableBody = document.getElementById('projectsTableBody');
  
  tableBody.innerHTML = projects.map(project => `
    <tr>
      <td>${project.title}</td>
      <td>${project.formsCount}</td>
      <td>${project.lastFormDate || 'N/A'}</td>
      <td>${formatDate(project.creative_at)}</td>
      <td>
        <button class="btn btn-info view-forms-btn" data-id="${project.id}">Просмотр</button>
      </td>
    </tr>
  `).join('');
  
  document.querySelectorAll('.view-forms-btn').forEach(btn => {
    btn.addEventListener('click', (e) => {
      const projectId = e.target.getAttribute('data-id');
      navigateTo(`/projects/${projectId}/forms`);
    });
  });
}