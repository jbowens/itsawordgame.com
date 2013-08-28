package model;

import java.util.Collection;

/**
 * A Trie factory for common patterns of constructing tries.
 *
 * @author jbowens
 */
public class TrieFactory
{

    /**
     * Constructs a Trie from a collection of strings.
     */
    public TrieNode constructFromCollection(Collection<String> strs)
    {
        TrieNode root = new TrieNode();
        
        for (String str : strs)
        {
            root.insertString(str);
        }

        return root;
    }

}
