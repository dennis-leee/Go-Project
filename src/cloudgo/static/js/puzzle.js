var blank = 15; //空白块初始化所在的位置
var startGame = false;
var ready = false;
var count = 0;
var time = 0;
var clock;


//结构初始化
window.onload = function() {
    var game = document.getElementById("game");
    var blockList = document.createElement("ul");
    blockList.className = "block_list";
    for(var num = 0; num < 16; num++) {
        var block = document.createElement("li");
        block.addEventListener("click", move);
        blockList.appendChild(block);
    }
    game.appendChild(blockList);
    initialization();
    document.getElementById("start").addEventListener("click", start);
    document.getElementById("start").addEventListener("click", timeCount);
    document.getElementById("resetting").addEventListener("click", reLayout);
}

//拼图初始化
function initialization() {
    var blockList = document.getElementsByTagName("li");
    for(var num = 0; num < 16; num++) {
        blockList[num].className = "puzzle_pieces correct_position_" + num;
        blockList[num].id = "current_position_" + num;
    }
    blockList[15].className ="blank";
}

//拼图布局
function reLayout() {
    if(startGame) {
        return;
    }
    initialization();
    var numOfHard = 50;   //从初始位置随机走50步
    for(var swapTime = 0; swapTime < numOfHard; swapTime++) {
        var blankBlock = document.getElementsByClassName("blank")[0];
        var num_pos = blankBlock.id.search("n_") + 2;
        var currentPosition = parseInt(blankBlock.id.slice(num_pos));
        var positionToBeMove = [];
        if (currentPosition != 3 && currentPosition != 7 && currentPosition != 11 && currentPosition != 15) {
        positionToBeMove.push(currentPosition + 1);
        }
        if (currentPosition != 0 && currentPosition != 4 && currentPosition != 8 && currentPosition != 12) {
        positionToBeMove.push(currentPosition - 1);
        }
        if (currentPosition != 0 && currentPosition != 1 && currentPosition != 2 && currentPosition != 3) {
        positionToBeMove.push(currentPosition - 4);
        }
        if (currentPosition != 12 && currentPosition != 13 && currentPosition != 14 && currentPosition != 15) {
        positionToBeMove.push(currentPosition + 4);
        }
        var numOfPosition = positionToBeMove.length;
        var random = Math.floor(numOfPosition * Math.random());
        swap(currentPosition, positionToBeMove[random]);
        blank = positionToBeMove[random];
    }
    ready = true;
}

//开始游戏
function start() {
    if(startGame) {
        initialization();
        clear();
        document.getElementById("start").textContent = "Start Game";
        document.getElementById("state").value = "Waiting";
    } else {
        document.getElementById("state").value = "Playing";
        document.getElementById("start").textContent = "Stop Game";
        if(!ready) {
            reLayout();
        }
        startGame = true;
    }
    document.getElementById("time").value = "0s";
    document.getElementById("steps").value = "0";
}

//时钟
function timeCount() {
    if(startGame) {
        time += 1;
        document.getElementById("time").value = time + "s";
        clock = setTimeout("timeCount()", 1000);
    }
}

//移动图片
function move(event) {
    if(startGame) {
        var num_pos = event.target.id.search("n_") + 2;
        var currentPosition = parseInt(event.target.id.slice(num_pos));
        if((currentPosition != 3 && currentPosition != 7 && currentPosition != 11 && currentPosition != 15 && currentPosition + 1 == blank) || (currentPosition != 0 && currentPosition != 4 && currentPosition != 8 && currentPosition != 12 && currentPosition - 1 == blank) ||  (currentPosition != 0 && currentPosition != 1 && currentPosition != 2 && currentPosition != 3 && currentPosition - 4 == blank) || (currentPosition != 12 && currentPosition != 13 && currentPosition != 14 && currentPosition != 15 && currentPosition + 4 == blank )) {
            swap(currentPosition, blank);
            blank = currentPosition;
            console.log(count);
            count++;
            document.getElementById("steps").value = count;
        }
    }
}

//交换方块
function swap(currentBlock, anotherBlock) {
    var blockA = document.getElementById("current_position_" + currentBlock);
    var blockB = document.getElementById("current_position_" + anotherBlock);
    blockA.id = "current_position_" + anotherBlock;
    anotherBlock = currentBlock;
    blockB.id = "current_position_" + anotherBlock;
    if(startGame) {
        check();
    }
}

//判断是否完成拼图
function check() {
    var win = true;
    var puzzleList = document.getElementsByClassName("puzzle_pieces");
    for(var num = 0; num < 15; num++) {
        var num_pos = puzzleList[num].id.search("n_") + 2;
        var currentPosition = parseInt(puzzleList[num].id.slice(num_pos));
        if(currentPosition != num) {
            win = false;
            break;
        }
    }
    if(win) {
        document.getElementById("state").value = "You win!";
        document.getElementById("start").textContent = "Start Game";
        clear();
    }
}

//重置各参数
function clear() {
    startGame = false;
    ready = false;
    clearTimeout(clock);
    blank = 15;
    count = 0;
    time = 0;
}
