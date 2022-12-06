import { useEffect, useState } from "react";
import api from "./api";
import debounce from "lodash.debounce";
import { TimePlot, createDataSet } from "./components/time_chart";


function App() {

  const [schemas, setSchemas] = useState(null);
  const kurants = [0.1, 0.5, 0.8, 1, 5, 10, 20, 50]

  const [n, setN] = useState(32.0);
  const [eps, setEps] = useState(0.03);
  const [task, setTask] = useState("1");
  const [schema, setschema] = useState("1");
  const [errors, setErrors] = useState(0);
  const [k, setK] = useState(1);

  const [originalGrid, setOriginalGrid] = useState([]);
  const [numerGrid, setNumerGrid] = useState([]);
  const [timePoints, setTimePoints] = useState([]);
  const [id, setId] = useState(null);

  const [numerSol, setNumerSol] = useState([]);
  const [original, setOriginal] = useState([]);

  const toParams = () => {
    return { n: parseInt(n), eps: parseFloat(eps), task: parseInt(task), schema: parseInt(schema), t, kurant: parseFloat(k) }
  }

  useEffect(() => {
    const func = async () => {
      var data = await api.getSchemas();
      if (data.status == 200) {
        setSchemas(data.data);
      };

      data = await api.getInitials(toParams())
      if (data.status == 200) {
        setOriginalGrid(data.data.xOriginal);
        setNumerGrid(data.data.xNumer);
        setTimePoints(toTimepoints(data.data.timePoints));
        setId(data.data.id);
      }

      data = await api.getSolution(toParams(), 0, data.data.id);
      if (data.status == 200) {
        setOriginal(data.data.original)
        setNumerSol(data.data.numerical)
        setDataSet([
          createDataSet(originalGrid, data.data.original, "original", "rgb(255,0,0)"),
          createDataSet(numerGrid, data.data.numerical, "numerical", "rgb(0,255,0)"),
        ])
      }
    }
    func()
    return () => { }
  }, []);

  const t = 2.0

  const toTimepoints = (floats) => {
    var res = []
    for (let index = 0; index < floats.length; index++) {
      res[index] = index + 1
    }
    return res
  }

  const debounceN = debounce(query => {
    if (!query) return setN(0)
    var parsed = parseFloat(query)
    if (parsed <= 0) parsed = 1;
    if (parsed > 10000) parsed = 10000;
    setN(parsed)
  }, 10)

  const debounceEps = debounce(query => {
    if (!query) return setEps(0)
    var parsed = parseFloat(query)
    if (parsed > 1) parsed = 1;
    if (parsed < 0) parsed = 0;
    setEps(parsed)
  }, 10)

  const [dataSet, setDataSet] = useState([]);

  const data = () => {
    if (originalGrid.length != original.length || numerGrid.length != numerSol.length) {
      return []
    }
    return [
      createDataSet(originalGrid, original, "original", "rgb(255,0,0)"),
      createDataSet(numerGrid, numerSol, "numerical", "rgb(0,255,0)"),
    ]
  }

  useEffect(() => {
    setDataSet(data())
  }, [numerSol, original]);

  useEffect(() => {
    const func = async () => {
      var data = await api.getInitials(toParams())
      if (data.status == 200) {
        setOriginalGrid(data.data.xOriginal);
        setNumerGrid(data.data.xNumer);
        setTimePoints(toTimepoints(data.data.timePoints));
        setId(data.data.id);
      }

      data = await api.getSolution(toParams(), 0, data.data.id);
      if (data.status == 200) {
        setOriginal(data.data.original)
        setNumerSol(data.data.numerical)
      }
    }
    func()
    setDataSet(data())
    return () => { }
  }, [n, eps, schema, k]);

  const newData = async (cur) => {
    var dat = await api.getSolution(toParams(), cur - 1, id);
    if (dat && dat.status != 200) {
      console.log(dat)
    }
    if (!dat) {
      return data()
    }
    return [
      createDataSet(originalGrid, dat.data.original, "original", "rgb(255,0,0)"),
      createDataSet(numerGrid, dat.data.numerical, "numerical", "rgb(0,255,0)"),
    ]
  }

  return (
    <div className="flex">
      <div className="flex mx-8 my-6 flex-1">
        <div className="flex flex-col p-4">
          <h2 className="text-3xl font-bold mb-2">Heat solver</h2>
          <p>Made by Talkanbaev Artur</p>
          <div className="mt-8">
            <h3 className="font-bold text-lg mb-2">Manual</h3>
            <p>1. Edit parameters - use the arrows or write the number down</p>
            <p>2. Use mouse and wheel to navigate the graph</p>
            <p>3. Hower over graphs to see values</p>
          </div>
          <h4 className="mt-4 text-lg font-bold">Parameters</h4>
          <div className="flex flex-col">
            <div className="flex flex-col my-2">
              <label className="text-gray-800-text-small mb-2">Grid size</label>
              <input
                type="number"
                className="outline-none ring hover:shadow-xl mr-auto ring-green-400 rounded-lg px-4"
                value={n}
                onInput={(e) => debounceN(e.target.value)}
              />
            </div>
            <div className="flex flex-col my-4">
              <label className="text-gray-800-text-small mb-2">Epsilon</label>
              <input
                type="number"
                className="outline-none ring hover:shadow-xl ring-green-400 rounded-lg px-4 mr-auto"
                value={eps}
                step=".01"
                onInput={(e) => debounceEps(e.target.value)}
              />
            </div>
            <div className="flex flex-col my-4">
              <label className="text-gray-800-text-small mb-2">Schema</label>
              <select className="outline-none ring ring-green-400 mr-auto pr-10 rounded-lg hover:shadow-xl bg-white py-1 pl-2" value={schema} onChange={(e) => { setschema(e.target.value) }}>
                {schemas ? schemas.map(s => {
                  return (
                    <option key={s.id} value={s.id}>{s.name}</option>
                  )
                }) : ""}
              </select>
            </div>
            <div className="flex flex-col my-4">
              <label className="text-gray-800-text-small mb-2">Kurant</label>
              <select className="outline-none ring ring-green-400 mr-auto pr-10 rounded-lg hover:shadow-xl bg-white py-1 pl-2" value={k} onChange={(e) => { setK(e.target.value) }}>
                {kurants ? kurants.map(s => {
                  return (
                    <option key={s} value={s}>{s}</option>
                  )
                }) : ""}
              </select>
            </div>
            <div className="flex flex-col my-4">
              <label className="text-gray-800-text-small mb-2">Task</label>
              <select className="outline-none ring ring-green-400 mr-auto pr-10 rounded-lg hover:shadow-xl bg-white py-1 pl-2" value={task} onChange={(e) => { setTask(e.target.value) }}>
                <option value="1">Task #1</option>
              </select>
            </div>
            <div className="flex flex-col my-6 space-y-4">
              <h4 className="font-bold text-lg">Error values</h4>
              <p>Error: {errors}</p>
            </div>
          </div>
        </div>
        <div className="flex flex-1 w-2/3">
          {
            TimePlot(dataSet, timePoints, newData)
          }
        </div>
      </div>
    </div>
  )
}

export default App


//{JSON.stringify(backendData)}
//
