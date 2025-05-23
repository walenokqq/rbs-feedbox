export function initRouter() {
  window.addEventListener('popstate', route);
  document.addEventListener('DOMContentLoaded', route);
}

export function route(event) {
  event && event.preventDefault();
  
  const path = window.location.pathname;
  const app = document.getElementById('app');
  
  app.innerHTML = '';
  
  if (path === '/' || path === '/projects') {
    import('../pages/projects').then(module => {
      module.renderProjectsPage();
    });
  } else if (path.startsWith('/projects/') && path.includes('/forms')) {
    const projectId = path.split('/')[2];
    import('../pages/projectForms').then(module => {
      module.renderProjectFormsPage(projectId);
    });
  } else if (path.startsWith('/feedback/')) {
    const formId = path.split('/')[2];
    import('../pages/feedback').then(module => {
      module.renderFeedbackPage(formId);
    });
  } else if (path.startsWith('/form/')) {
    const formId = path.split('/')[2];
    import('../pages/form').then(module => {
      module.renderFormPage(formId);
    });
  } else {
    import('../pages/notFound').then(module => {
      module.renderNotFoundPage();
    });
  }
}


export function navigateTo(path) {
  window.history.pushState({}, '', path);
  const popStateEvent = new PopStateEvent('popstate');
  window.dispatchEvent(popStateEvent);
}