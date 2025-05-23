const API_URL = 'http://localhost:8080/api/';


async function sendRequest({ endpoint, method = 'GET', body = null }) {
  const url = `${API_URL}${endpoint}`;
  const header = {
    'Content-Type': 'application/json',
  };

  const config = {
    method,
    header,
    body: body ? JSON.stringify(body) : null
  };

  try {
    const response = await fetch(url, config);
    
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    return await response.json();
  } catch (error) {
    console.error('API request failed:', error);
    throw error;
  }
}


export async function getProjects() {
  // return [
  //   { id: '1', title: 'Проект №1', formsCount: 3, lastFormDate: '2025-05-20', projectDate:"2025-01-15" },
  //   { id: '2', title: 'Проект №2', formsCount: 5, lastFormDate: '2025-05-18', projectDate:"2025-03-25" }
  // ];
  const data = sendRequest({ endpoint: 'projects' });

  return data;
}

export async function createProject(title) {
  return { id: Date.now().toString(), title, formsCount: 0, lastFormDate: '' };
}


export async function getProjectForms(projectId) {
  return sendRequest({ endpoint: `projects/${projectId}/forms` });
}



export async function createForm(projectId, title) {
  return sendRequest({
    endpoint: `projects/${projectId}/forms`,
    method: 'POST',
    body: { title }
  });
}

export async function getFormFeedback(formId) {
  return sendRequest({ endpoint: `forms/${formId}/responses` });

  
}

export async function updateFeedbackStatus(feedbackId, status) {
  return { success: true };
}

export async function getForm(formId) {
  return sendRequest({
    endpoint: `form/${formId}`
  });
}
export async function saveFormResponse(formId, data) {
  return sendRequest({
    endpoint: `feedback/${formId}`,
    method: 'POST',
    body: data
  });
}