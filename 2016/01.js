// ============================

let canvas = document.getElementById("can");
let ctx = canvas.getContext("2d");

// ============================

let programX = "R1, R3, L2, L5, L2, L1, R3, L4, R2, L2, L4, R2, L1, R1, L2, R3, L1, L4, R2, L5, R3, R4, L1, R2, L1, R3, L4, R5, L4, L5, R5, L3, R2, L3, L3, R1, R3, L4, R2, R5, L4, R1, L1, L1, R5, L2, R1, L2, R188, L5, L3, R5, R1, L2, L4, R3, R5, L3, R3, R45, L4, R4, R72, R2, R3, L1, R1, L1, L1, R192, L1, L1, L1, L4, R1, L2, L5, L3, R5, L3, R3, L4, L3, R1, R4, L2, R2, R3, L5, R3, L1, R1, R4, L2, L3, R1, R3, L4, L3, L4, L2, L2, R1, R3, L5, L1, R4, R2, L4, L1, R3, R3, R1, L5, L2, R4, R4, R2, R1, R5, R5, L4, L1, R5, R3, R4, R5, R3, L1, L2, L4, R1, R4, R5, L2, L3, R4, L4, R2, L2, L4, L2, R5, R1, R4, R3, R5, L4, L4, L5, L5, R3, R4, L1, L3, R2, L2, R1, L3, L5, R5, R5, R3, L4, L2, R4, R5, R1, R4, L3";
let test1 = "R2, L3";
let test2 = "R2, R2, R2";
let test3 = "R5, L5, R5, R3";
let test4 = "R8, R4, R4, R8";
let test5 = "L2, R3";
let test6 = "R1, R3, L2, L5, L2, L1, R3, L4, R2, L2, L4, R2, L1, R1, L2, R3";

function Pos(x, y, token) {
	this.x = x;
	this.y = y;
	this.token = token;
	this.distance = () => this.x + this.y;
	this.toString = () => `${this.x},${this.y}`
}

function Taxicab(program) {
	this.locations = [
		new Pos(20, 20)
	];
	this.rotations = [];
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
		let startX = this.locations[0].x,
			startY = this.locations[0].y;
		let endX = this.locations[this.locations.length - 1].x,
			endY = this.locations[this.locations.length - 1].y;
		let distance = Math.abs(startX - endX) + Math.abs(startY - endY);
		return distance;
	}

	this.run = function () {
		let facing = 0;
		let x = 0,
			y = 0;
		for (let i = 0; i < this.tokens.length; i++) {
			let token = this.tokens[i];
			let [turn, num] = this.parseToken(token);
			if (turn === "R")
				facing++;
			else
				facing--;
			facing = facing % 4;
			console.log(`facing ${facing}`)

			for (let j = 0; j < num; j++) {
				if (facing === 0) {
					console.log("north");
					y++;
				} else if (facing === 1) {
					console.log("east");
					x++;
				} else if (facing === 2) {
					console.log("south");
					y--;
				} else if (facing === 3) {
					console.log("west");
					x--;
				}
				let pos = new Pos(x+20, y+20, token);
				this.locations.push(pos);
				if (this.seen.hasOwnProperty(pos.toString())) {
					this.seen[pos.toString()]++
				} else
					this.seen[pos.toString()] = 0
			}

		}
		Object.keys(this.seen).forEach(key => {
			let [x, y] = key.split(",");
			if (this.seen[key] === 1) console.log(`seen ${key} ${Math.abs(x) + Math.abs(y)}`)
		});
		return this;
	}
}
let scale = 20;
var a = new Taxicab(test6).run();
try {
	ctx.save();
	ctx.lineWidth = 1;
	ctx.translate(can.width / 2, can.height / 2);
	a.locations.forEach((point, idx) => {
		let previous = idx === 0 ? a.locations[0] : a.locations[idx - 1];
		ctx.strokeStyle = `rgb(${(idx * scale ) % 255}, 128, 128)`
		ctx.save();
		ctx.beginPath();
		ctx.moveTo(previous.x * scale, previous.y * scale);
		ctx.lineTo(point.x * scale, point.y * scale);

		ctx.stroke();
		ctx.restore();
	});
/* 
	Object.keys(a.seen).forEach(key => {
		let [x, y] = key.split(",");
		if (a.seen[key] === 1) {
			ctx.save();
			ctx.fillRect(x * scale, y * scale, 10, 10);
			ctx.restore();

		}

	}); */
	ctx.restore();
} catch (e) {
	console.error(e)
}
//a.show();

/* new Taxicab(test2).run();
new Taxicab(test3).run();
new Taxicab(test4).run().findCrossSection();
new Taxicab(programX).run().findCrossSection(); */