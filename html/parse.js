async function getOrders() {
    const response = await fetch('http://localhost:8080/api/order/list');
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
      orderCard.className = 'orderCard';
      orderCard.innerHTML = `# Заказа: ${id}<br> Блюдо: ${dishName}<br> Порций: ${quantity}<br> Готовность: ${isReady}<br> Выдан: ${isServed} <br>
        <br>
        <button onclick="deleteOrder(${id})" class="btn btn-outline-danger">Удалить</button>
        <button onclick="updateOrder(${id})" class="btn btn-outline-primary">Обновить</button>`;
      orderContainer.appendChild(orderCard);
    });
  }
  
  async function deleteOrder(orderId) {
    await fetch(`http://localhost:8080/api/order/delete`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ "ID": orderId }),
    });
    getOrders();
  }
  
  async function createOrderForm() {
    const newDish = document.getElementById("dishName").value;
    const newQuantity = document.getElementById("dishCount").value;
    
    await fetch('http://localhost:8080/api/order/create', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ "Dish": newDish, "Quantity": parseInt(newQuantity), "IsReady": false, "IsServed": false }),
    });
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