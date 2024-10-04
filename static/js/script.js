document.addEventListener("DOMContentLoaded", function () {
    // Timeline interactivity
    const timelineEvents = document.querySelectorAll(".timeline-event");
    timelineEvents.forEach(event => {
        event.addEventListener("click", function () {
            alert(event.querySelector("h3").textContent);
        });
    });

    // Simple map implementation (without external libraries)
    const map = document.getElementById("map");
    const battles = [
        {{range .Battles}}
        {
            name: "{{.Name}}",
            location: "{{.Location}}",
            lat: {{.Latitude}},
            lon: {{.Longitude}},
            description: "{{.Description}}",
            date: "{{.Date}}"
        },
        {{end}}
    ];

    battles.forEach(battle => {
        const marker = document.createElement("div");
        marker.className = "marker";
        marker.style.position = "absolute";
        marker.style.top = (battle.lat * 5) + "px";  // Simplified mapping
        marker.style.left = (battle.lon * 5) + "px";
        marker.title = battle.name + " (" + battle.date + "): " + battle.description;
        map.appendChild(marker);
    });
});
