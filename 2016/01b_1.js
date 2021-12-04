let programX = "R1, R3, L2, L5, L2, L1, R3, L4, R2, L2, L4, R2, L1, R1, L2, R3, L1, L4, R2, L5, R3, R4, L1, R2, L1, R3, L4, R5, L4, L5, R5, L3, R2, L3, L3, R1, R3, L4, R2, R5, L4, R1, L1, L1, R5, L2, R1, L2, R188, L5, L3, R5, R1, L2, L4, R3, R5, L3, R3, R45, L4, R4, R72, R2, R3, L1, R1, L1, L1, R192, L1, L1, L1, L4, R1, L2, L5, L3, R5, L3, R3, L4, L3, R1, R4, L2, R2, R3, L5, R3, L1, R1, R4, L2, L3, R1, R3, L4, L3, L4, L2, L2, R1, R3, L5, L1, R4, R2, L4, L1, R3, R3, R1, L5, L2, R4, R4, R2, R1, R5, R5, L4, L1, R5, R3, R4, R5, R3, L1, L2, L4, R1, R4, R5, L2, L3, R4, L4, R2, L2, L4, L2, R5, R1, R4, R3, R5, L4, L4, L5, L5, R3, R4, L1, L3, R2, L2, R1, L3, L5, R5, R5, R3, L4, L2, R4, R5, R1, R4, L3";
let test1 = "R2, L3";
let test2 = "R2, R2, R2";
let test3 = "R5, L5, R5, R3";
let test4 = "R8, R4, R4, R8";

function Pos(x, y) {
    this.x = x;
    this.y = y;
}

function Taxicab(program) {
    this.history = [
        new Pos(0, 0)
    ];
    this.seen = {};
    this.program = program;
    this.tokens = program.replace(/\s+/g, "").split(",");
    this.parseDigit = function (token) {
        let result = [];
        for (let i = 0; i < token.length; i++) {
            let char = token.charAt(i);
            if (char >= "0" && char <= "9") {
                result.push(parseInt(char));
            }
        }
        return parseInt(result.join(""));
    };
    this.parseToken = function (token) {
        let result = [];
        for (let i = 0; i < token.length; i++) {
            let char = token.charAt(i);
            if (char >= "0" && char <= "9") {
                let number = this.parseDigit(token.substring(i, token.length))
                result.push(number);
                i += number.length;
            }
            if (char == "L" || char == "R") {
                result.push(char);
            }
        }
        return result;
    };
    this.manhatten = function manhatten() {

        let startX = this.history[0].x,
            startY = this.history[0].y;
        let endX = this.history[this.history.length - 1].x,
            endY = this.history[this.history.length - 1].y;
        let distance = Math.abs(startX - endX) + Math.abs(startY - endY);
        return distance;

    }

    this.run = function () {
      
    }

    this.findCrossSection = function () {
        let seen = {};
        let pos;
        for(let i = 0; i < this.history.length; i++) {
            pos = this.history[i];
            let s = `${pos.x};${pos.y}`;
            if (seen.hasOwnProperty(s)) {
                console.log("first time crossed", JSON.stringify(pos));
                break;
            } else {
                seen[s] = 1;
            }
        }
        let start = this.history[0];
        console.log("Distance ", Math.abs(start.x - pos.x) + Math.abs(start.y - pos.y));
    }
}

var a = new Taxicab(test1).run();
console.log(`${JSON.stringify(a)}`);
/* 
new Taxicab(test2).run();
new Taxicab(test3).run();
new Taxicab(test4).run().findCrossSection();
new Taxicab(programX).run().findCrossSection();
 */
// console.log( distance);

/* 
let intersect = a.reduce((obj, pos) => {
    let p = pos.join(";");
    if (obj.hasOwnProperty(p)) {
        obj[p] += 1;
    } else {
        obj[p] = 1;
    }
    return obj;
}, {});

let indexes = [];
let h = a.forEach((elm, idx, arr) => {
    let j =  elm.join(";");
    if (intersect[j] > 1) {
        indexes.push(idx);
    }
}, []);

console.log(indexes);

function manhatten(arr) {
    let start = arr[0];
    let end = arr[arr.length - 1];
    console.log(`start ${start}, end ${end}`);
    let distance = Math.abs(start[0] - end[0]) + Math.abs(start[1] - end[1]);
    console.log(`distance ${distance}`)
}

let b = [[0,0], a[h]];
console.log(manhatten(b));
//console.log(`intersect ${JSON.stringify(intersect, null, 4)}`); */