// [@POS [X,Y]]
pos = [
    [482, 84], // CRAFT END

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
    [78, 252],
    [132, 252],
    [186, 252],
    [240, 252],
    [294, 252],
    [348, 252],
    [402, 252],
    [456, 252],
    [24, 360], // THIRD ROW
    [78, 252],
    [132, 252],
    [186, 252],
    [240, 252],
    [294, 252],
    [348, 252],
    [402, 252],
    [456, 252],

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

/**
 *
 * @param {Array.<{Amount: number, Type: string}|null>} entries
 * @param {{Amount: number, Type: string}|null} onCursor
 */
function process(entries, onCursor) {
    let V = jQuery("#view");
    V.empty();
    entries.forEach(function (v, i) {
        if (v) {
            V.append(`<div class="item" style="top: ${pos[i][1]}px; left: ${pos[i][0]}px; background-image: url(img/views/${v.Type}.png)" data-type="${v.Type}"><div class='tl'>${i}</div><div class='br'>${v.Amount}</div></div>`)
        }
    });
    if (onCursor) {
        V.append(`<div class="item" style="top: ${cursorPos[1]}px; left: ${cursorPos[0]}px; background-image: url(img/views/${onCursor.Type}.png)" data-type="${onCursor.Type}"><div class='tl'>-1</div><div class='br'>${onCursor.Amount}</div></div>`)
    }
}
