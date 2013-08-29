package model;

import java.util.Collection;

/**
 * Stores information about a user scoring event. This class is often
 * marshalled into json to send to the user over the websocket.
 *
 * @author jbowens
 */
public class ScoreEvent
{

    /**
     * Words scored in this event.
     */
    public final Collection<Word> words;

    public ScoreEvent(Collection<Word> wordsScored)
    {
        words = wordsScored;
    }

}
