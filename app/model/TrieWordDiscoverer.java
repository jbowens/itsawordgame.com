package model;

import java.util.ArrayList;
import java.util.Collection;

/**
 * A WordDiscoverer implementation that uses a Trie to find words.
 */
public class TrieWordDiscoverer implements WordDiscoverer
{

    // Possible nodes that still match the current suffix
    protected Collection<TrieNode> m_previousNodes;

    public TrieWordDiscoverer(TrieNode root)
    {
        m_previousNodes = new ArrayList<TrieNode>();
        m_previousNodes.add(root);
    }

    public Collection<Word> processLetter(char letter)
    {
        Collection<Word> words = new ArrayList<Word>();

        for (TrieNode node : m_previousNodes)
        {
            TrieNode replacementNode = node.getChildForCharacter(letter);
            m_previousNodes.remove(replacementNode);
            if (replacementNode != null)
            {
                // Add this node back in, it has a subtrie
                m_previousNodes.add(replacementNode);

                if(replacementNode.isWord())
                {
                    // This is a matching word! Return this a scored word.
                    String str = replacementNode.getAncestralString();
                    Word word = new Word(str);
                    words.add(word);
                }
            }
        }

        return words;
    }

}
