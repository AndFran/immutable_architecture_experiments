# immutable_architecture_experiments
Just some implementations of immutable architecture from the book: The Art of Immutable Architecture by Michael L. Perry


## immutable_tree: 

      is a simple implementation of the immutable tree structure shown on page 12 of the second edition.
      Currently, it just supports inserting new nodes. A new root node is returned with the pointers
      of the internal data structure changing just for those nodes that need to change, that is a path
      down to where we insert the new node.

      given the tree:

<pre>
                             12
			/        \
                      7             27
                   /    \        /      \
                 3       9      17       32
                /  \          /    \
               1    5        14     25
</pre>	
 
if we want to insert 100 the path i.e. the nodes that would change are: 12, 27, 32 with 100 being inserted on the new 
node of 32:

<pre>
                             12'
			/        \
                      7             27'
                   /    \        /      \
                 3       9      17       32'
                /  \          /    \        \
               1    5        14     25      100
</pre>	

All other nodes that do not need to change stay the same i.e. at the same address.


 


