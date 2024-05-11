async function getOrders() {
  const response = await fetch('/api/order/list');
  const data = await response.json();
  console.log(data);
  displayOrders(data);
}

function displayOrders(orders) {
  const orderContainer = document.getElementById('orderContainer');
  orderContainer.innerHTML = '';

  orders.forEach(order => {
    let id = order.id
    let dishName = order.dish || "Блюдо не указано";
    let quantity = order.quantity || 0;
    let isReady = order.isready ? "Да" : "Нет";
    let isServed = order.isserved ? "Да" : "Нет";
    const orderCard = document.createElement('div');
    orderCard.className = 'card';
    orderCard.style.width = "16rem";
    orderCard.innerHTML = `<div class="card-body">
      <h6 class="card-title"><b>Номер заказа: ${id}</b></h6>
      <p class="card-text">Блюдо: ${dishName}</p>
      <p class="card-text">Кол-во порций: ${quantity}</p>
      <p class="card-text">Готов: ${isReady}</p>
      <p class="card-text">Выдан: ${isServed}</p>
      <button onclick="deleteOrder(${id})" class="btn btn-outline-danger">Удалить</button>
      <button onclick="updateOrder(${id})" class="btn btn-outline-primary">Обновить</button>
    </div>
  `;
    orderContainer.appendChild(orderCard);
  });
}

async function deleteOrder(orderId) {
  await fetch(`/api/order/delete`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ "ID": orderId }),
  });
  getOrders();
  appendAlert(`Заказ #${orderId} успешно удален`, "success");
}

async function createOrderForm() {
  const newDish = document.getElementById("dishName");
  const newQuantity = document.getElementById("dishCount");

  await fetch('/api/order/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ "Dish": newDish.value, "Quantity": parseInt(newQuantity.value), "IsReady": false, "IsServed": false }),
  });

  newDish.value = "";
  newQuantity.value = "";
  getOrders();
}

async function updateOrder(orderId) {
  await fetch('/api/order/update', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ "ID": orderId, "IsReady": true }),
  });
  getOrders();
}

getOrders();


const alertPlaceholder = document.getElementById('liveAlertPlaceholder');
const appendAlert = (message, type) => {
  const wrapper = document.createElement('div')
  wrapper.innerHTML = [
    `<div class="alert alert-${type} alert-dismissible" role="alert">`,
    `   <div>${message}</div>`,
    '   <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>',
    '</div>'
  ].join('')

  alertPlaceholder.append(wrapper)
}

