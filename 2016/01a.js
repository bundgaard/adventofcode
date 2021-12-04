let degrees = 0;
let x = 0, y = 0;
let program = "R5, L5, R5, R3";
let puzzle = "R1, R3, L2, L5, L2, L1, R3, L4, R2, L2, L4, R2, L1, R1, L2, R3, L1, L4, R2, L5, R3, R4, L1, R2, L1, R3, L4, R5, L4, L5, R5, L3, R2, L3, L3, R1, R3, L4, R2, R5, L4, R1, L1, L1, R5, L2, R1, L2, R188, L5, L3, R5, R1, L2, L4, R3, R5, L3, R3, R45, L4, R4, R72, R2, R3, L1, R1, L1, L1, R192, L1, L1, L1, L4, R1, L2, L5, L3, R5, L3, R3, L4, L3, R1, R4, L2, R2, R3, L5, R3, L1, R1, R4, L2, L3, R1, R3, L4, L3, L4, L2, L2, R1, R3, L5, L1, R4, R2, L4, L1, R3, R3, R1, L5, L2, R4, R4, R2, R1, R5, R5, L4, L1, R5, R3, R4, R5, R3, L1, L2, L4, R1, R4, R5, L2, L3, R4, L4, R2, L2, L4, L2, R5, R1, R4, R3, R5, L4, L4, L5, L5, R3, R4, L1, L3, R2, L2, R1, L3, L5, R5, R5, R3, L4, L2, R4, R5, R1, R4, L3"
let tokens = puzzle.replace(/\s/gi, "").split(",");
for(let i = 0; i < tokens.length; i++) {
    let token = tokens[i];
    degrees += token.charAt(0) === "R" ? 90 : -90;
    let steps = parseInt(token.charAt(1));
    switch(degrees % 360) {
        case 0:
        y += steps;
        break;
        case 90:
        case -270:
        x += steps;
        break;
        case 180:
        case -180:
        y -= steps;
        break;
        case 270:
        case -90:
        x -= steps;
        break;
    }
}

let away = Math.abs(x) + Math.abs(y);
console.log("Your destination ", away, "blocks away");