(function() {

    // Setup
    $(document).ready(function(e) {

        initGame();

        // TODO: Do setup
    });

    var width = 5;
    var height = 4;

    function initGame() {
        $.get('/init-game', function(round) {
            round = $.parseJSON(round);
            var cells = round.board.cells;

            // Sort the cells by row, column
            cells.sort(function (a, b) { return a.row * round.board.width + a.column - (b.row * round.board.width + b.column)});

            var clearfix = $("#board .clearfix");

            // Empty out the board
            clearfix.siblings().remove();
    
            // Add the new cells
            for (var i = 0; i < cells.length; i++)
            {
                var cell = constructCell(cells[i].letter, cells[i].row, cells[i].column);
                $(cell).insertBefore(clearfix); 
            }
        });
    }

    function constructCell(letter, row, column) {
        var cell = $('<div class="cell"></div>');
        cell.data('row', row);
        cell.data('column', column);

        var innerCell = $('<div class="inner-cell"></div>');
        var letterSpan = $('<span class="letter"></span>');
        letterSpan.text(letter);
        innerCell.append(letterSpan);
        cell.append(innerCell);
        letterSpan.hover(cellHover, cellExit);
        return cell;
    }

    function cellHover(e) {
        var cell = $(this).closest('.cell');
        cell.addClass('hover');
    }

    function cellExit(e) {
        var cell = $(this).closest('.cell');
        var innerCell = cell.find('.inner-cell');
        cell.removeClass('hover');

        // Do an exit fade animation
        if (cell.data('animationTimeout')) {
            clearTimeout(cell.data('animationTimeout'));
            cell.data('animationTimeout', null);
        }

        var num = 30;
        function animationTick() {
            var alpha = num / 30.0;
            var color = 'rgba(163,67,99,'+alpha+')';
            innerCell.css('background-color', color);
            num--;

            if (num > 0)
                cell.data('animationTimeout', setTimeout(animationTick, 30));
        }

        animationTick();

    }

})();
