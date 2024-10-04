document.addEventListener("DOMContentLoaded", function() {
    const tabsContainer = document.getElementById('tabs');
    const biographyElement = document.getElementById('biography');
    const characterImage = document.getElementById('characterImage');
    const characterBio = document.getElementById('characterBio');
    const anecdotesSection = document.getElementById('anecdotesSection');
    const referencesSection = document.getElementById('referencesSection');
    const biographyTitle = document.getElementById('biographyTitle');
    const anecdotesList = document.getElementById('anecdotesList');
    const referencesList = document.getElementById('referencesList');

    // Cache les sections "Anecdotes et Références" par défaut
    anecdotesSection.style.display = "none";
    referencesSection.style.display = "none";

    // Fonction pour charger les biographies depuis l'API
    function loadBiographies() {
        fetch('/api/biographies')
            .then(response => response.json())
            .then(data => {
                // Créer les onglets pour chaque personnage
                data.forEach(bio => {
                    const tab = document.createElement('div');
                    tab.className = 'tab';
                    tab.dataset.name = bio.name;
                    tab.textContent = bio.name;

                    tab.addEventListener('click', function() {
                        // Mettre à jour le titre de la biographie
                        biographyTitle.textContent = bio.name;
                        
                        // Afficher la biographie, anecdotes et références
                        characterImage.src = `path/to/images/${bio.name.replace(/ /g, '_')}.jpg`; // Modifiez selon votre structure d'image
                        characterBio.textContent = bio.biography;

                        // Gérer les anecdotes
                        anecdotesList.innerHTML = '';
                        if (bio.anecdotes) {
                            anecdotesSection.style.display = "block";
                            bio.anecdotes.split(';').forEach(anecdote => {
                                const li = document.createElement('li');
                                li.textContent = anecdote;
                                anecdotesList.appendChild(li);
                            });
                        } else {
                            anecdotesSection.style.display = "none";
                        }

                        // Gérer les références
                        referencesList.innerHTML = '';
                        if (bio.references) {
                            referencesSection.style.display = "block";
                            bio.references.split(';').forEach(reference => {
                                const li = document.createElement('li');
                                li.textContent = reference;
                                referencesList.appendChild(li);
                            });
                        } else {
                            referencesSection.style.display = "none";
                        }

                        // Mettre à jour l'état des onglets
                        const activeTab = tabsContainer.querySelector('.tab.active');
                        if (activeTab) {
                            activeTab.classList.remove('active');
                        }
                        tab.classList.add('active');
                    });

                    tabsContainer.appendChild(tab);
                });
            })
            .catch(error => console.error('Erreur lors du chargement des biographies:', error));
    }

    loadBiographies(); // Charge les biographies au chargement de la page
});
