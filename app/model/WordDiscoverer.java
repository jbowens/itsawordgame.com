package model;

import java.util.Collection;

/**
 * Word discovers take a sequence of letters, letter by letter, and return
 * any words contained within the sequence, incrementally.
 *
 * @author jbowens
 */
public interface WordDiscoverer
{

    /**
     * Adds the given letter to the current sequence and then returns a collection of
     * all words discovered in the letter sequence including the newly added letter.
     */
    public Collection<Word> processLetter(char letter);

}
