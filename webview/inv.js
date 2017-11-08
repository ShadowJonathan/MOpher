Array.prototype.clone = function () {
    return this.slice(0);
};

// [@POS [X,Y]]
invPos = [
    [462, 84], // CRAFT END

    [294, 54], // CRAFT
    [348, 54],
    [294, 108],
    [348, 108],

    [24, 24], // ARMOR
    [24, 78],
    [24, 132],
    [24, 186],

    [24, 252], // INVENTORY, FIRST ROW
    [78, 252],
    [132, 252],
    [186, 252],
    [240, 252],
    [294, 252],
    [348, 252],
    [402, 252],
    [456, 252],
    [24, 306], // SECOND ROW
    [78, 306],
    [132, 306],
    [186, 306],
    [240, 306],
    [294, 306],
    [348, 306],
    [402, 306],
    [456, 306],
    [24, 360], // THIRD ROW
    [78, 360],
    [132, 360],
    [186, 360],
    [240, 360],
    [294, 360],
    [348, 360],
    [402, 360],
    [456, 360],

    [24, 426], // HOTBAR
    [78, 426],
    [132, 426],
    [186, 426],
    [240, 426],
    [294, 426],
    [348, 426],
    [402, 426],
    [456, 426],

    [231, 186] // SECOND HAND
];

cursorPos = [126, 153];

craftingPos = [
    [371, 105],

    [90, 51],
    [144, 51],
    [198, 51],

    [90, 105],
    [144, 105],
    [198, 105],

    [90, 159],
    [144, 159],
    [198, 159],
];


function generateInv(offset) {
    let x = offset[0];
    let y = offset[1];
    return [
        [x, y], // INVENTORY, FIRST ROW
        [x + 54, y],
        [x + (54 * 2), y],
        [x + (54 * 3), y],
        [x + (54 * 4), y],
        [x + (54 * 5), y],
        [x + (54 * 6), y],
        [x + (54 * 7), y],
        [x + (54 * 8), y],
        [x, y + 54], // SECOND ROW
        [x + 54, y + 54],
        [x + (54 * 2), y + 54],
        [x + (54 * 3), y + 54],
        [x + (54 * 4), y + 54],
        [x + (54 * 5), y + 54],
        [x + (54 * 6), y + 54],
        [x + (54 * 7), y + 54],
        [x + (54 * 8), y + 54],
        [x, y + 108], // THIRD ROW
        [x + 54, y + 108],
        [x + (54 * 2), y + 108],
        [x + (54 * 3), y + 108],
        [x + (54 * 4), y + 108],
        [x + (54 * 5), y + 108],
        [x + (54 * 6), y + 108],
        [x + (54 * 7), y + 108],
        [x + (54 * 8), y + 108],
        [x, y + 174], // HOTBAR
        [x + 54, y + 174],
        [x + (54 * 2), y + 174],
        [x + (54 * 3), y + 174],
        [x + (54 * 4), y + 174],
        [x + (54 * 5), y + 174],
        [x + (54 * 6), y + 174],
        [x + (54 * 7), y + 174],
        [x + (54 * 8), y + 174],
    ]
}

// type: X, Y
typeIndex = {
    chest: [7, 39],
    cobblestone: [10, 27],
    dirt: [13, 30],
    sand: [3, 31],
    mossy_cobblestone: [25, 30],
    gold_ore: [24, 31],
    gravel: [21, 30],
    red_sand: [1, 31],
};

/**
 * @param {number[][]} pos
 * @param {number[]} cpos
 */
function initPos(pos, cpos) {
    let V = $("#view");
    V.empty();
    for (let i in pos) {
        if (i == "clone")
            continue;
        let p = pos[i];
        V.append(`<div class="item" style="top: ${p[1]}px; left: ${p[0]}px;" id="${i}"><div class='tl'>${i}</div><div class='br hidden'></div></div>`)
        let I = $(`#${i}`);
        I.on("click", () => {
            sock.send(i + "")
        });
        I.on("contextmenu", e => {
            e.preventDefault();
            sock.send("* " + i + "")
        })
    }
    V.append(`<div class="item" style="top: ${cpos[1]}px; left: ${cpos[0]}px;" id="-1"><div class='tl'>-1</div><div class='br hidden'></div></div>`)
    let I = $(`#-1`);
    I.on("click", () => {
        sock.send("-1")
    });
    I.on("contextmenu", e => {
        e.preventDefault();
        sock.send("* -1")
    })
}

var currentType = null;
var currentAmount = null;

/**
 *
 * @param {Array.<{Amount: number, Type: string}|null>} entries
 * @param {{Amount: number, Type: string}|null} onCursor
 * @param {number} type
 */
function process(entries, onCursor, type) {
    let V = $(".item");
    // noinspection EqualityComparisonWithCoercionJS
    if (type != currentType || entries.length != currentAmount) {
        let v = $("#view");
        v.css("height", "");
        v.css("width", "");
        switch (type) {
            case -1:
                v.css("background-image", "url(img/inventory.png)");
                initPos(invPos, cursorPos);
                break;
            case 0:
            case 1:
                v.css("background-image", `url(img/chest_${(entries.length - 36) / 9}.png`);
                v.css("height", 342 + (((entries.length - 36) / 9) * 54));
                v.css("width", 528);
                initPos(generateChest(entries.length), [0, 0]);
                break;
            case 2:
                v.css("background-image", `url(img/crafting_table.png`);
                v.css("height", 498);
                v.css("width", 528);
                initPos(craftingPos.clone().concat(generateInv([24, 252])), [23, 105]);
                break;
        }
        currentType = type;
        currentAmount = entries.length;
    }

    V.each((_, e) => {
        e = $(e);
        let i = parseInt(e.attr('id'));
        let br = e.find(".br");
        if (i >= 0) {
            if (entries[i]) {
                let v = entries[i];
                doImage(v.Type, e);
                br.html(v.Amount);
                br.toggleClass("hidden", false);
            } else {
                e.css("background-image", "");
                e.css("background-position", "");
                br.toggleClass("hidden", true)
                e.toggleClass("all-icon", false)
            }
        } else {
            if (onCursor) {
                let v = onCursor;
                doImage(v.Type, e);
                br.toggleClass("hidden", false);
                br.html(v.Amount)
            } else {
                e.css("background-image", "");
                e.css("background-position", "");
                br.toggleClass("hidden", true);
                e.toggleClass("all-icon", false)
            }
        }
    })
}

/**
 * @param {number} amount
 * @returns {number[][]}
 */
function generateChest(amount) {
    var pos = [24, 54];
    var invStartOffset = 96;
    var invOffsetToLast = invStartOffset - 54;

    if ((amount - 36) % 9) {
        window.alert("ERROR " + amount);
        throw new Error("ERROR " + amount);
    }

    let rows = (amount - 36) / 9;

    let allPos = [];

    for (let i = 0; i < rows; i++) {
        let bx = pos[0];
        for (let i = 0; i < 9; i++) {
            allPos.push(pos.clone())
            pos[0] += 54;
        }
        pos[0] = bx;
        pos[1] += 54
    }

    pos[1] += invOffsetToLast;

    return allPos.concat(generateInv(pos));
}


function doImage(type, e) {
    if (!typeIndex[type]) {
        e.css("background-image", `url(img/views/${type}.png`);
        e.css("background-position", "");
        e.toggleClass("all-icon", false)
    } else {
        e.css("background-image", `url(img/all.png`);
        e.css("background-position", `-${typeIndex[type][0] * 48}px -${typeIndex[type][1] * 48}px`);
        e.toggleClass("all-icon", true)
    }
}