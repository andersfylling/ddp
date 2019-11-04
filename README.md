# Discord docs parser (ddp)
Converts all the types, consts, flags, etc. into csv files.

Goal:
 - Make the titles consistent:
   - use name, value, type, type-hash, description, + whatever the rest discord uses; to properly yield insight into the data.    
     - Name is a trivial word
     - Value is the actual value
     - Type can be int, uint, snowflake, etc. 
     - Description is just a description of whatever
 - separate data into constants and structures.
 - allow name and value to be the same, by setting the title to "name/value"
 - add new columns to give a relationship overview:
   - Type-hash would be a generated hash to uniquely identify other discord data structures (eg. Message)
