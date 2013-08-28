package model;

import java.util.Collection;

/**
 * A WordDiscoverer implementation that uses a Trie to find words.
 */
public class TrieWordDiscoverer implements WordDiscoverer
{

    /**
     * Constructs a word discoverer that uses a Trie to back it from a dictionary of words. In
     * actual use, the dictionary of words should be only the words that actually appear within
     * the board, not all words in the English dictionary. Performance will suffer if you use
     * all words in the dictionary, because construction of the trie will take a while.
     */
    public TrieWordDiscoverer(Collection<String> words)
    {

    }

    public Collection<Word> processLetter(char letter)
    {
        // TODO: Implement
        return null;
    }

}
