# Runs
A bug bounty auxiliary tool that traverses the content of each line of the file and fills it into the command for execution.

run:
```
Usage: runs <textfile> <command>

Iterate through each line of the text file and execute it in the command.

Parameters:
        <textfile>
                The path to the text file that will be processed. This file should exist on your system and be readable.
        <command>
                The command to execute after processing the text file. This must be a recognized command that the program can perform.

Example:
        runs domain.txt "assetfinder -subs-only {*}"

        Iterate through each line in domain.txt, fill it into "{*}", and then execute.

Notes:
  - Ensure that the path to the text file is correct and the file is accessible.
  - The command should be a valid operation that this program is designed to handle.
  - Review the available commands and their syntax if necessary.

version:
        0.0.1 - 面包狗
```
