export function formatDate(dateString) {
  if (!dateString) return 'N/A';
  
  const date = new Date(dateString);
  if (isNaN(date.getTime())) return 'Invalid Date';
  
  const day = date.getDate();
  const month = date.toLocaleString('ru-RU', { 
    month: 'long',
    day: 'numeric'
  }).split(' ')[1];
  
  const year = date.getFullYear();
  
  return `${day} ${month} ${year} года`;
}