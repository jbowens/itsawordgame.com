(function() {
  // Constants
  var ANIMATION_GRANULARITY = 30;
  var ANIMATION_PAUSE_MILLI = 40;

  // Global variables
  var width = 5;
  var height = 4;
  var ws = null;
  var game = null;

  $(document).ready(function(e) {
    connect();
    // TODO: Do additional setup.
  });

  /* Receives any game stream messages from the server. These are mostly
   * reports of successful word scores.
   */
  function receiveEvent(event) {
    var data = JSON.parse(event.data);

    console.log(data);

    // Handle errors
    if (data.error) {
      // TODO: Handle gracefully
      console.log(data.error.message);
      return;
    }

    if (data.message_type == "announce_game") {
      game = data.game;
      setupBoard(game);
    }

  }

  /* Event handler for socket closes.
   */
  function onSocketClose(event) {
    ws = null;
    console.log(event);
  }

  /* Event handler for socket errors.
   */
  function onSocketError(event) {
    // TODO: Handle gracefully
    console.log(event);
  }

  /* Opens a Web Socket connection to the server on which user events
   * and game statuses can stream back and forth. There must not be
   * an existing connection. The resulting connection is saved in
   * the ws global.
   */
  function connect() {
    if (ws) {
      throw new Error("A game socket is already open.");
    }

    var WS = !!window.MozWebSocket ? MozWebSocket : WebSocket;
    ws = new WS("ws://" + window.location.host + "/connect"); ws.onmessage = receiveEvent;
    ws.onclose = onSocketClose;
    ws.onerror = onSocketError;
  }

  /*******************************************************************
   *                         UI Functionality                        *
   *******************************************************************/

  /* Updates the board UI to represent the given game.
   */
  function setupBoard(g) {
    var cells = g.board.cells;
    var width = g.board.width;
    var height = g.board.height;

    var clearfix = $("#board .clearfix");

    // Empty out the board
    clearfix.siblings().remove();

    // Add the new cells
    for (var i = 0; i < cells.length; i++) {
      var cell = constructCell(String.fromCharCode(cells[i].letter), i / width, i % height, cells[i].id);
      $(cell).insertBefore(clearfix);
    }
  }

  /* Constructs a cell with the given data and returns the outer
   * div of the cell. Used in refreshing the board when a new game
   * begins.
   */
  function constructCell(letter, row, column, id) {
    var cell = $('<div class="cell"></div>');
    cell.data('id', id);
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

  /* Event listener for when the user begins hovering over a cell.
   */
  function cellHover(e) {
    var cell = $(this).closest('.cell');
    cell.addClass('hover');
  }

  /* Event listener for when the user's cursor leaves a cell.
   */
  function cellExit(e) {
    var cell = $(this).closest('.cell');
    var innerCell = cell.find('.inner-cell');
    cell.removeClass('hover');

    // Fade from the hover color to the regular cell background
    fade(innerCell, {r: 163, g:67, b:99});
  }

  /* Animates a color change for the given element, from the given color to
   * transparency.
   */
  function fade(el, color) {
    el = $(el);
    // If there's an existing animation, cancel it. This newer
    // one takes precedence.
    if (el.data('animationTimeout')) {
      clearTimeout(el.data('animationTimeout'));
      el.data('animationTimeout', null);
    }

    var animationFrame = ANIMATION_GRANULARITY;
    function animationTick() {
      var alpha = animationFrame / ANIMATION_GRANULARITY;
      var cssColor = 'rgba('+color.r+','+color.g+','+color.b+','+alpha+')';
      el.css('background-color', cssColor);
      animationFrame--;

      if (animationFrame > 0) {
        el.data('animationTimeout', setTimeout(animationTick, ANIMATION_PAUSE_MILLI));
      }
    }
    animationTick();
  }
})();
