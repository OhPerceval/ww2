document.addEventListener("DOMContentLoaded", function() {
    loadTechnologies();
    loadStrategies();
});

function loadTechnologies() {
    fetch('/api/technologies')
        .then(response => response.json())
        .then(data => {
            const technologiesList = document.getElementById('technologiesList');
            data.forEach(tech => {
                const card = document.createElement('div');
                card.className = 'card';
                card.innerHTML = `
                    <img src="${tech.image}" alt="${tech.name}">
                    <h3>${tech.name}</h3>
                    <p>${tech.description}</p>
                `;
                technologiesList.appendChild(card);
            });
        })
        .catch(error => console.error('Erreur lors du chargement des technologies:', error));
}

function loadStrategies() {
    fetch('/api/strategies')
        .then(response => response.json())
        .then(data => {
            const strategiesList = document.getElementById('strategiesList');
            data.forEach(strategy => {
                const card = document.createElement('div');
                card.className = 'card';
                card.innerHTML = `
                    <h3>${strategy.name}</h3>
                    <p>${strategy.description}</p>
                `;
                strategiesList.appendChild(card);
            });
        })
        .catch(error => console.error('Erreur lors du chargement des stratégies:', error));
}

function showSection(section) {
    const sections = document.querySelectorAll('.tab-content');
    sections.forEach(sec => {
        sec.classList.remove('active');
    });

    const buttons = document.querySelectorAll('.tab-button');
    buttons.forEach(btn => {
        btn.classList.remove('active');
    });

    document.getElementById(section).classList.add('active');
    document.querySelector(`.tab-button[onclick="showSection('${section}')"]`).classList.add('active');
}
