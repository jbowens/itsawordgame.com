package controllers;

import model.*;
import org.codehaus.jackson.*;
import org.codehaus.jackson.map.ObjectMapper;
import play.*;
import play.mvc.*;
import views.html.*;

public class Application extends Controller
{

    /* Dimensions for the default board. */
    protected static final int BOARD_WIDTH = 5;
    protected static final int BOARD_HEIGHT = 4;

    /* The round generator to use for generating new rounds. */
    protected static RoundGenerator m_roundGenerator = new RoundGenerator(BOARD_WIDTH, BOARD_HEIGHT);
  
    public static Result index()
    {
        return ok(index.render());
    }

    /**
     * Called by the client to initialize a new game. It returns the arrangement for the client's
     * next round. In the future this will use precomputed data, but for testing purposes this
     * currently computes an entire new game from scratch.
     */
    public static Result initGame()
    {
        Round round = m_roundGenerator.generateRandomRound();

        JsonNode json = new ObjectMapper().valueToTree(round);

        return ok(json.toString());
    }

    /**
     * A 404 catchall.
     * TODO: Actually route this
     */
    public static Result fourOhFour()
    {
        return notFound(fourohfour.render());
    }

    public static WebSocket<JsonNode> game()
    {
        return new WebSocket<JsonNode>() {
            // Called when the WebSocket handshake is done
            public void onReady(WebSocket.In<JsonNode> in, WebSocket.Out<JsonNode> out)
            {
                // TODO: Implement
            }
        };
    }
  
}
