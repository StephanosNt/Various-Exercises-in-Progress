document.addEventListener("DOMContentLoaded", () => {
    fetch("/api/artists")
        .then((response) => response.json())
        .then((data) => {
            const artistList = document.getElementById("artist-list");
            data.forEach((artist) => {
                const div = document.createElement("div");
                div.innerHTML = `
                    <h2>${artist.name}</h2>
                    <img src="${artist.image}" alt="${artist.name}" width="200">
                    <a href="artist.html?id=${artist.id}">View Details</a>
                `;
                artistList.appendChild(div);
            });
        });
});
