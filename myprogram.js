let movies = [
    {
        "name": "Титаник",
        "genre": "Драма",
        "rating": 8.4
    },
    {
        "name": "Выживший",
        "genre": "Боевик",
        "rating": 7.8
    },
    {
        "name": "Начало",
        "genre": "Научная фантастика",
        "rating": 8.7
    },
    {
        "name": "Волк с Уолл-Стрит",
        "genre": "Комедия",
        "rating": 8.8
    },
    {
        "name": "Однажды в Голливуде",
        "genre": "Комедия",
        "rating": 7.6
    }
];

function sortByTitle(movies) {
    movies.sort((a, b) => {
        let titleA = a.name.toLowerCase();
        let titleB = b.name.toLowerCase();
        if (titleA < titleB) return -1;
        if (titleA > titleB) return 1;
        return 0;
    });
}

function sortByRating(movies) {
    movies.sort((a, b) => {
        return b.rating - a.rating;
    });
}

function sortByGenre(movies) {
    movies.sort((a, b) => {
        let genreA = a.genre.toLowerCase();
        let genreB = b.genre.toLowerCase();
        if (genreA < genreB) return -1;
        if (genreA > genreB) return 1;
        return 0;
    });
}

function calculateAverageRating(movies) {
    let sum = 0;
    for (let i = 0; i < movies.length; i++) {
        sum += movies[i].rating;
    }
    return sum / movies.length;
}

sortByTitle(movies);
console.log("Sorted by title:", movies);

sortByRating(movies);
console.log("Sorted by rating:", movies);

sortByGenre(movies);
console.log("Sorted by genre:", movies);

let moviesDiv = document.getElementById("movies");
for (let i = 0; i < movies.length; i++) {
    let movie = movies[i];
    let movieDiv = document.createElement("div");
    movieDiv.innerHTML = `<h3>${movie.name}</h3><p>Genre: ${movie.genre}, Rating: ${movie.rating}</p>`;
    moviesDiv.appendChild(movieDiv);
}

let averageRating = calculateAverageRating(movies);
let averageRatingDiv = document.getElementById("average-rating");
averageRatingDiv.innerHTML = `<h3>Средний рейтинг:</h3><p>${averageRating}</p>`;

