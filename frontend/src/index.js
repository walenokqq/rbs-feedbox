import './styles/main.css';
import { initRouter } from './utils/router';
import { renderProjectsPage } from './pages/projects';

// Инициализация роутера
initRouter();

// Загрузка начальной страницы
window.addEventListener('load', () => {
  const path = window.location.pathname;
  if (path === '/' || path === '/projects') {
    renderProjectsPage();
  }
});