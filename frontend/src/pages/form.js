import { navigateTo } from '../utils/router';
import { getForm, saveFormResponse } from '../utils/api';

export function renderFormPage(formId) {
  const app = document.getElementById('app');
  
  app.innerHTML = `
    <div class="container">
      <div class="form-header">
        <button id="backToFormsBtn" class="btn btn-secondary">← Назад к формам</button>
        <h2 id="formTitle">Загрузка формы...</h2>
      </div>
      
      <div id="formContainer">
        <div id="formFieldsContainer"></div>
        <button id="submitFormBtn" class="btn btn-primary" disabled>Отправить</button>
      </div>
      
      <div id="formMessage" class="mt-3"></div>
    </div>
  `;

  // Навигация назад
  document.getElementById('backToFormsBtn').addEventListener('click', () => {
    navigateTo(`/projects/${projectId}/forms`);
  });

  loadAndRenderForm(formId);
}

async function loadAndRenderForm(formId) {
  try {
    const formData = await getForm(formId);
    document.getElementById('formTitle').textContent = formData.title;
    
    // Парсим JSON schema формы
    const schema = JSON.parse(formData.schema);
    renderFormFields(schema);
    
    // Активируем кнопку отправки
    document.getElementById('submitFormBtn').disabled = false;
    document.getElementById('submitFormBtn').addEventListener('click', () => {
      submitForm(formId, schema);
    });
    
  } catch (error) {
    console.error('Ошибка при загрузке формы:', error);
    document.getElementById('formContainer').innerHTML = `
      <div class="alert alert-danger">Не удалось загрузить форму.</div>
    `;
  }
}

function renderFormFields(schema) {
  const container = document.getElementById('formFieldsContainer');
  container.innerHTML = '';
  
  schema.fields.forEach((field, index) => {
    const fieldElement = createFormField(field, index);
    container.appendChild(fieldElement);
  });
}

function createFormField(field, index) {
  const fieldGroup = document.createElement('div');
  fieldGroup.className = 'form-group mb-3';
  
  const label = document.createElement('label');
  label.htmlFor = `field-${index}`;
  label.textContent = field.label;
  if (field.required) {
    label.innerHTML += ' <span class="text-danger">*</span>';
  }
  
  let input;
  
  switch (field.type) {
    case 'text':
    case 'email':
    case 'number':
      input = document.createElement('input');
      input.type = field.type;
      input.className = 'form-control';
      input.id = `field-${index}`;
      input.name = field.name;
      input.required = field.required || false;
      if (field.placeholder) {
        input.placeholder = field.placeholder;
      }
      break;
      
    case 'textarea':
      input = document.createElement('textarea');
      input.className = 'form-control';
      input.id = `field-${index}`;
      input.name = field.name;
      input.required = field.required || false;
      if (field.placeholder) {
        input.placeholder = field.placeholder;
      }
      if (field.rows) {
        input.rows = field.rows;
      }
      break;
      
    case 'select':
      input = document.createElement('select');
      input.className = 'form-control';
      input.id = `field-${index}`;
      input.name = field.name;
      input.required = field.required || false;
      
      field.options.forEach(option => {
        const optionElement = document.createElement('option');
        optionElement.value = option.value;
        optionElement.textContent = option.label;
        input.appendChild(optionElement);
      });
      break;
      
    case 'checkbox':
      input = document.createElement('input');
      input.type = 'checkbox';
      input.id = `field-${index}`;
      input.name = field.name;
      input.className = 'form-check-input';
      break;
      
    case 'radio':
      fieldGroup.className = 'mb-3';
      label.className = 'd-block';
      label.textContent = field.label;
      
      field.options.forEach((option, optIndex) => {
        const radioGroup = document.createElement('div');
        radioGroup.className = 'form-check';
        
        const radioInput = document.createElement('input');
        radioInput.type = 'radio';
        radioInput.id = `field-${index}-${optIndex}`;
        radioInput.name = field.name;
        radioInput.value = option.value;
        radioInput.className = 'form-check-input';
        radioInput.required = field.required || false;
        
        const radioLabel = document.createElement('label');
        radioLabel.htmlFor = `field-${index}-${optIndex}`;
        radioLabel.className = 'form-check-label';
        radioLabel.textContent = option.label;
        
        radioGroup.appendChild(radioInput);
        radioGroup.appendChild(radioLabel);
        fieldGroup.appendChild(radioGroup);
      });
      return fieldGroup;
      
    default:
      input = document.createElement('input');
      input.type = 'text';
      input.className = 'form-control';
      input.id = `field-${index}`;
      input.name = field.name;
  }
  
  if (field.type !== 'radio') {
    fieldGroup.appendChild(label);
    fieldGroup.appendChild(input);
  }
  
  if (field.description) {
    const helpText = document.createElement('small');
    helpText.className = 'form-text text-muted';
    helpText.textContent = field.description;
    fieldGroup.appendChild(helpText);
  }
  
  return fieldGroup;
}

async function submitForm(formId, schema) {
  const formMessage = document.getElementById('formMessage');
  formMessage.innerHTML = '';
  formMessage.className = 'mt-3';
  
  try {
    // Собираем данные формы
    const formData = {};
    schema.fields.forEach((field, index) => {
      if (field.type === 'checkbox') {
        formData[field.name] = document.getElementById(`field-${index}`).checked;
      } else if (field.type === 'radio') {
        const selectedRadio = document.querySelector(`input[name="${field.name}"]:checked`);
        formData[field.name] = selectedRadio ? selectedRadio.value : null;
      } else {
        formData[field.name] = document.getElementById(`field-${index}`).value;
      }
    });
    
    // Отправляем данные на бэкенд
    const response = await saveFormResponse(formId, formData);
    
    // Показываем сообщение об успехе
    formMessage.className = 'mt-3 alert alert-success';
    formMessage.textContent = 'Форма успешно отправлена!';
    
    // Очищаем форму (кроме radio и checkbox)
    schema.fields.forEach((field, index) => {
      if (field.type !== 'checkbox' && field.type !== 'radio') {
        document.getElementById(`field-${index}`).value = '';
      } else if (field.type === 'checkbox') {
        document.getElementById(`field-${index}`).checked = false;
      }
    });
    
  } catch (error) {
    console.error('Ошибка при отправке формы:', error);
    formMessage.className = 'mt-3 alert alert-danger';
    formMessage.textContent = 'Произошла ошибка при отправке формы. Пожалуйста, попробуйте ещё раз.';
  }
}