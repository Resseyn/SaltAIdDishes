document.addEventListener ('DOMContentLoaded', function() {
    let elems = document.querySelectorAll('select');
    let instances = M.FormSelect.init(elems)
    let buttons = document.querySelectorAll('.select-btn-group button');
    let generateButton = document.getElementById('generate');

    buttons.forEach(function(button) {
        console.log("123")
        button.addEventListener('click', function() {
            // Убираем активный класс у всех кнопок
            buttons.forEach(function(btn) {
                btn.classList.remove('active');
            });

            // Добавляем активный класс только для нажатой кнопки
            button.classList.add('active');
        });
    });
})

document.getElementById('generate').addEventListener('click', function () {
    let selectedButton = document.querySelector('.btn-group button.active');
    let selectedValue = ""
    if (selectedButton) {
        selectedValue = selectedButton.textContent;
    }

    let kitchenElement = document.getElementById('kitchen');
    let kitchenValue = kitchenElement.options[kitchenElement.selectedIndex].value;
    let typeElement = document.getElementById('type');
    let typeValue = typeElement.options[typeElement.selectedIndex].value;
    let priceElement = document.getElementById('price');
    let priceValue = priceElement.options[priceElement.selectedIndex].value;
    let speciesElement = document.getElementById('species');
    let speciesValue = speciesElement.options[speciesElement.selectedIndex].value;
    let consistElement = document.getElementById('consist');
    let consistValue = consistElement.options[consistElement.selectedIndex].value;

    let data = {"params":[kitchenValue, selectedValue, typeValue, priceValue, speciesValue, consistValue]}

    // Выводим выбранное значение в консоль (или можете использовать по своему усмотрению)
    console.log('Selected Value:', data);
    const resultRow = document.getElementById('dishRow');
    if (resultRow.style.display != "none") {
        let toDelete = resultRow.firstElementChild
        toDelete.remove()
    }
    fetch('http://localhost:8080/generate', {
        method: "POST",
        headers: {'Content-Type': "application/json"},
        body: JSON.stringify(data)
    })
        .then(response => response.json()) // Предполагается, что сервер возвращает JSON
        .then(data => {
            // Обработка полученных данных

            const resultDiv = document.createElement('div');
            if (data.english == undefined) {
                resultDiv.innerHTML = `
            <div>ошибочка</div>
            <img src="https://i.pinimg.com/564x/37/60/d3/3760d344114afbc05a1266f1f985d382.jpg" alt="https://i.pinimg.com/564x/37/60/d3/3760d344114afbc05a1266f1f985d382.jpg">
            <div>вышла блин(((</div>
            `;
            } else {
                resultDiv.innerHTML = `
            <div><h5 style="text-align: center">${data.english}</h5></div>
            <img src="${data.image}" alt="${data.image}" style="margin-left: auto;
    margin-right: auto;
    display: block;">
            <div><p style="text-align: center">${data.recipe}</p></div>
            `;
            }

            // resultDiv.textContent = JSON.stringify(data, null, 2); // Преобразование данных в строку и вывод
            resultRow.appendChild(resultDiv);
            resultRow.style.display = "flex"
        })
        .catch(error => {
            console.error('Произошла ошибка при запросе:', error);
        });
});
