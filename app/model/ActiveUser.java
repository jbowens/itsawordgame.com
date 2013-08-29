package model;

import org.codehaus.jackson.*;
import play.mvc.*;

/**
 * Represents an active user. A user is considered active if they're currently
 * connected via a WebSocket. This means at the very least they have the page
 * open. They may not be interacting with the game at all.
 *
 * @author jbowens
 */
public class ActiveUser
{
    /* The state of this user. */
    protected UserState m_state = UserState.WAITING_FOR_USER;

    /* In and out web sockets. */
    protected WebSocket.In<JsonNode> m_inSocket;
    protected WebSocket.Out<JsonNode> m_outSocket;

    /**
     * Constructor for an active user. Takes in the in and out sockets of the
     * Web Socket connection.
     */
    public ActiveUser(WebSocket.In<JsonNode> in, WebSocket.Out<JsonNode> out)
    {
        m_inSocket = in;
        m_outSocket = out;
    }

    /**
     * Enum for possible user states
     */
    public static enum UserState
    {
        /**
         * The user is choosing what to do.
         */
        WAITING_FOR_USER,

        /**
         * The user is waiting for a new round. The server should send this
         * user a round as soon as one is available.
         */
        WAITING_FOR_GAME,

        /**
         * The user is currently playing a game.
         */
        PLAYING
    }
}
