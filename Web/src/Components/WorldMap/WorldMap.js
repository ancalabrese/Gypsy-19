import React, { useEffect } from 'react'
import CountryListDummy from '../../Data/DummyLists';
import { Card } from 'react-bootstrap';
import Style from './WorldMap.module.css'
import Theme from '../../index.scss'

const WorldMap = (props) => {
    const mapData = [['Country', 'Travel List']]
    const RED = "Red";
    const AMBER = "Amber";
    const GREEN = "Green";
    const ListMapping = {
        [RED]: 0,
        [AMBER]: 1,
        [GREEN]: 2,
    }
    mapData.push(["United Kingdom", ListMapping[GREEN]])
    mapData.push(["Ireland", ListMapping[GREEN]])

    for (const [key, values] of Object.entries(CountryListDummy)) {
        var k = "";
        switch (key) {
            case "red":
                k = RED;
                break;
            case "amber":
                k = AMBER;
                break;
            case "green":
                k = GREEN;
                break;
        }
        values.forEach(e => {
            mapData.push([e.name, ListMapping[k]])
        });
    }

    useEffect(() => {
        window.google.charts.setOnLoadCallback(renderMap);
    })

    const disableMapLegend = () => {
        var legend = document.querySelector("#map-container > div > div:nth-child(1) > div > svg > g > g:nth-child(3)")
        if (legend != null) {
            legend.setAttribute("style", "display:none");
        }
    }

    const mapReadyHandler = () => {
        disableMapLegend();
    }

    const renderMap = () => {
        const redColor = getComputedStyle(document.documentElement).getPropertyValue("--red-list-color").slice(1);
        const amberColor = getComputedStyle(document.documentElement).getPropertyValue("--amber-list-color").slice(1);
        const greenColor = getComputedStyle(document.documentElement).getPropertyValue("--green-list-color").slice(1);
        const mapBackground = getComputedStyle(document.documentElement).getPropertyValue("--primary-dark").slice(1);
        const strokeColor = getComputedStyle(document.documentElement).getPropertyValue("--primary-light").slice(1);
        var options = {
            colorAxis: {
                values: [ListMapping[RED], ListMapping[AMBER], ListMapping[GREEN]],
                colors: [redColor, amberColor, greenColor]
            },
            backgroundColor: mapBackground,
            datalessRegionColor: '#ffff',
            defaultColor: '#f5f5f5',
            stroke: strokeColor,
            strokeWidth: 5,
            domain: 'UK',
            tooltip:{
                trigger:"selection"
            }
        };
        const map = new window.google.visualization.GeoChart(document.getElementById("map-container"));
        const data = window.google.visualization.arrayToDataTable(mapData);
        window.google.visualization.events.addListener(map, 'ready', mapReadyHandler());
        map.draw(data, options);
    }

    return (
        <Card className={[Style['card'], "primary-dark-element"].join(" ")}>
            <Card.Body className={Style['card-body']}>
                <div id="map-container" />
            </Card.Body>
        </Card>
    )
}

export default WorldMap;