import React, { useEffect } from 'react'


const WorldMap = (props) => {
    const mapData = [
        ['Country', 'Travel List'],
        ['Germany', 'Grenn'],
        ['United States', 'Green'],
        ['Brazil', 'Red'],
        ['Canada', 'Red'],
        ['France', 'Red'],
        ['RU', 'Red']
    ];
        useEffect(() => {
            window.google.charts.setOnLoadCallback(renderMap);
        })

    const renderMap = () => {
        const options = {};
        const map = new window.google.visualization.GeoChart(document.getElementById("map-container"));
        const data = window.google.visualization.arrayToDataTable(mapData);
        map.draw(data, options);
    }

    return (
        <div id="map-container" className="min-vh-100">
            Map placeholder
        </div>
    )
}

export default WorldMap;