import { Line } from "react-chartjs-2"
import { Chart, PointElement, registerables } from 'chart.js';
import zoomPlugin from 'chartjs-plugin-zoom';
import { FaPause, FaPlay } from "react-icons/fa"
import { TbPlayerPause, TbPlayerPlay, TbPlayerSkipBack, TbPlayerSkipForward, TbPlayerTrackNext, TbPlayerTrackPrev } from "react-icons/tb"
import { useEffect, useState } from "react";
import { useCallback } from "react";
Chart.register(...registerables);
Chart.register(zoomPlugin);

//creates a data set
function createDataSet(xVals, yVals, label, color) {
    var dataset = { label: label, borderColor: color }
    if (xVals.length != yVals.length) {
        throw "invalid xy vals supplied: lengths do not match"
    }
    var data = []
    for (let i = 0; i < xVals.length; i++) {
        data[i] = { x: xVals[i], y: yVals[i] }
    }
    dataset.data = data;
    return dataset;
}

function PlotPoint(pointName, isCur, pointUpdateFunc) {
    var classN = "py-1 px-3 rounded-lg ring-1 ring-green-500 " + (isCur ? "bg-green-200" : "")
    return (
        <button key={pointName} className={classN} onClick={pointUpdateFunc}>
            {pointName}
        </button>
    )
}

//supply an array of sets created by createDataSet function
function TimePlot(initialDatasets, timePoints, updateDataFunc) {
    var options = {
        elements: {
            point: {
                radius: 0
            },
            line: {
                cubicInterpolationMode: "monotonic",
            }
        },
        responsive: true,
        scales: {
            x: {
                type: 'linear'
            },
            y: {
                type: 'linear'
            }
        },
        plugins: {
            zoom: {
                zoom: {
                    wheel: {
                        enabled: true,
                    },
                    mode: 'xy',
                },
                pan: {
                    enabled: true,
                    mode: "xy",
                }
            }
        }
    }

    const [play, setPlay] = useState(false);
    const [inter, setInter] = useState(null);
    const [cur, setCur] = useState(1);
    const [data, setData] = useState({ datasets: initialDatasets });

    useEffect(() => {
        setData({ datasets: updateDataFunc(cur) });
    }, [cur])

    const upCur = useCallback(
        () => {
            setCur((val) => {
                var min = Math.min(...timePoints),
                    max = Math.max(...timePoints);
                val += 1;
                if (val < min) {
                    return (min);
                } else if (val > max) {
                    setCur(max);
                } else {
                    setCur(val);
                }
            })
        },
        [cur],
    );

    useEffect(() => {
        if (!play && inter) {
            clearInterval(inter)
            setInter(null);
        }

        if (play && !inter) {
            var interv = setInterval(upCur, 100)
            setInter(interv);
        }
    }, [cur, play]);


    const updateCur = (val) => {
        var min = Math.min(...timePoints),
            max = Math.max(...timePoints);
        if (val < min) {
            setCur(min);
        } else if (val > max) {
            setCur(max);
        } else {
            setCur(val);
        }
    }

    return (
        <div>
            <Line data={data} options={options}></Line>
            <div className="flex flex-col">
                <div name="controls" className="my-4 mx-auto flex space-x-4">
                    <button className="ctrl-button" onClick={() => { updateCur(-1) }}>{TbPlayerSkipBack()}</button>
                    <button className="ctrl-button" onClick={() => { updateCur(cur - 1) }}>{TbPlayerTrackPrev()}</button>
                    <button className="ctrl-button" onClick={() => { setPlay(true) }}>{TbPlayerPlay()}</button>
                    <button className="ctrl-button" onClick={() => { setPlay(false) }}>{TbPlayerPause()}</button>
                    <button className="ctrl-button" onClick={() => { updateCur(cur + 1) }}>{TbPlayerTrackNext()}</button>
                    <button className="ctrl-button" onClick={() => { updateCur(Infinity) }}>{TbPlayerSkipForward()}</button>
                </div>
                {cur}
                <div name="points" className="flex mx-auto space-x-2 items-center">
                    <span>Time Points:</span>
                    {timePoints.map(e => (
                        PlotPoint(e, e == cur, () => { setCur(e) })
                    ))}
                </div>
            </div>
        </div>
    )

}

export { TimePlot, createDataSet };