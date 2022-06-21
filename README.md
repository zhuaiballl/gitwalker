# gitwalker
Output all versions of a local git repo, which could be used as test data for your research program.

### Notice
This program is under development. Current version just calls ```git checkout HEAD^``` once and once again, so only older versions in your current git branch could be generated. And if you want to switch back to your current version, you need to do that manually.

### Usage
1. Build ```gitwalker```.
```shell
cd your/gitwalker/directory
go build
```
2. Copy the compiled file ```gitwalker``` to the git repo that you want to process.
3. Run ```gitwalker``` there.
```shell
./gitwalker
```
4. Ignore the error messages. 😛
5. Output files are in the directory ```$HOME/.gitwalker/```