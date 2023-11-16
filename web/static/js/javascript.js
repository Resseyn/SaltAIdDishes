document.addEventListener ('DOMContentLoaded', function() {
    let elems = document.querySelectorAll('select');
    let instances = M.FormSelect.init(elems)
})

document.getElementById('generate').addEventListener('click', function () {
    // Отправка запроса на сервер
    const resultRow = document.getElementById('dishRow');
    if (resultRow.style.display != "none") {
        let toDelete = resultRow.firstElementChild
        toDelete.remove()
    }
    data = {"params":["Итальянская","Закуска"]}
    fetch('http://localhost:8080/generate', {
        method: "POST",
        headers: {'Content-Type': "application/json"},
        body: JSON.stringify(data)
    })
        .then(response => response.json()) // Предполагается, что сервер возвращает JSON
        .then(data => {
            // Обработка полученных данных
            const resultDiv = document.createElement('div');

            resultDiv.innerHTML = `
            <div>${data.english}</div>
            <img src="${data.image}" alt="${data.image}">
            `;
            // resultDiv.textContent = JSON.stringify(data, null, 2); // Преобразование данных в строку и вывод
            resultRow.appendChild(resultDiv);
            resultRow.style.display = "flex"
        })
        .catch(error => {
            console.error('Произошла ошибка при запросе:', error);
        });
});