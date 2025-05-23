export function renderNotFoundPage() {
  const app = document.getElementById('app');
  
  app.innerHTML = `
    <div class="container">
      <h1>404 - Page Not Found</h1>
      <p>The page you are looking for does not exist.</p>
      <button id="goHomeBtn" class="btn btn-primary">Go to Home</button>
    </div>
  `;
  
  document.getElementById('goHomeBtn').addEventListener('click', () => {
    window.history.pushState({}, '', '/projects');
    const popStateEvent = new PopStateEvent('popstate');
    window.dispatchEvent(popStateEvent);
  });
}