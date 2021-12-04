const express = require("express");
const app = express();
const fs = require("fs");

function traverse(path) {
    console.log(`traverse ${path}`);
    try {
        const entries = fs.readdirSync(path);
        let folders = entries.reduce((arr, name) => {
            if (/\d{4}/.test(name)) {
                arr.push(`${path}${name}`);
            }
            return arr;
        }, []);

        folders.forEach((/* arr, */ folder) => {
            let items = fs.readdirSync(folder).map(item => `${folder}/${item}`);
            console.log(items);

        },[]);

       // console.log(`${JSON.stringify(files)}`);
    } catch (err) {
        console.error(err);
    }
}
traverse("./");
app.use("/", express.static("./2016"));
app.listen(3000, () => console.log("running on port 3000"));