function initChart(xy) {

    new Chart("myChart", {
        type: "scatter",
        data: {
            datasets: [{
                pointRadius: 4,
                pointBackgroundColor: "rgb(0,0,255)",
                data: xy
            }]
        },
        options: {
            legend: { display: false },
            scales: {
                xAxes: [{ ticks: { min: 40, max: 160 } }],
                yAxes: [{ ticks: { min: 6, max: 16 } }],
            }
        }
    });
}

function getData() {
    var xyValues = [
        { x: 50, y: 7 },
        { x: 60, y: 8 },
        { x: 70, y: 8 },
        { x: 80, y: 9 },
        { x: 90, y: 9 },
        { x: 100, y: 9 },
        { x: 110, y: 10 },
        { x: 120, y: 11 },
        { x: 130, y: 14 },
        { x: 140, y: 14 },
        { x: 150, y: 15 }
    ];

    return xyValues;
}

var datasets = []
function getFiisChartData() {
    var oujp = [
        93.59, 92.7, 94.69, 94.17, 94.39
    ]

    return oujp
}

function initFiiChart(fiiValues, date) {
    new Chart("myFiiChart", {
        type: "line",
        data: {
            labels: date,
            datasets: createDatasets()
        },
        options: {
            legend: { display: true, },
            scales: {
                yAxes: [{ ticks: { min: 60, max: 110 } }],
            }
        }
    });
}

function getDateArray() {
    return ["28-04-2023T00:00:00",
        "02-05-2023T00:00:00",
        "03-05-2023T00:00:00",
        "04-05-2023T00:00:00",
        "05-05-2023T00:00:00"]
}

function createDatasets() {
    var oujp = [
        93.59, 92.7, 94.69, 94.17, 94.39
    ]
    var gtwr = [
        79.99, 78.41, 78.84, 79.16, 79.17
    ]


    var datasets = []
    datasets.push(
        {
            data: oujp,
            borderColor: "red",
            fill: false,
            label: "OUJP"
        }
    )
    datasets.push(
        {
            data: gtwr,
            borderColor: "blue",
            fill: false,
            label: "GTWR"
        }
    )
    return datasets
}

function extractFiiValesToArray(fiiHistoricData) {
    var fiiValues = []

    fiiHistoricData.array.forEach(element => {
        fiiValues.push(element.valor)
    });

    return fiiValues
}

document.addEventListener('alpine:init', () => {

})

function initAll() {
    Alpine.store('integer', 0)

    Alpine.store('arrayTest', {
        items: ["HGLG11"],

        addFII(fiiName) {
            this.items.push(fiiName)
        },

        getItems() {
            return items
        }
    })
}

function addFiiToScope(scopedFiis, newFii) {
    if (newFii != undefined && newFii != '') {
        scopedFiis.push(newFii)
    }
}

function removeFiiFromScope(scopedFiis, fiiIndex) {
    scopedFiis.splice(fiiIndex, 1)
}

function clearcontent(elementId) {
    element = document.getElementById(elementId)
    while (element != null) {
        element.outerHTML = "";
        element = document.getElementById(elementId)
    }


}