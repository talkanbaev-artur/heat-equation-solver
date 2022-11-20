import { useEffect, useState } from "react";
import api from "./api";
import { TimePlot, createDataSet } from "./components/time_chart";


function App() {

  const [backendData, setBackendData] = useState(null);

  useEffect(() => {
    const func = async () => {
      var data = await api.getTypes();
      if (data.status == 200) {
        setBackendData(data.data);
      };
    }
    func()
    return () => { }
  }, []);


  const getData = (c) => {
    var x = [];
    var y = [];
    for (let i = 0; i < 100; i++) {
      x[i] = i * 0.1;
      y[i] = x[i] ** Math.E + c;
    }
    return [x, y]
  }

  const [x, y] = getData(1)


  var data = [createDataSet(x, y, "original", "rgb(255,10,10)")]
  return (
    <div className="">
      <div className="w-3/4">
        {JSON.stringify(backendData)}
        {TimePlot(data, [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24], (cur) => {
          const [x, y] = getData(cur);
          return [createDataSet(x, y, "original", "rgb(255,10,10)")]
        })}
      </div>
    </div>
  )
}

export default App
