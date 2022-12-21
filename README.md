# gitwalker
Output all versions of a local git repo, which could be used as test data for your research program.

### Notice
This program is under development. If you want to switch back to your current version, you need to do that manually.

### Usage
1. Build ```gitwalker```.
```shell
cd your/gitwalker/directory
go install
```
2. Add `$GOPATH/bin` to your terminal's `PATH` if you have not.
3. `cd` to the git repo that you want to process.
4. Run ```gitwalker``` there.
```shell
gitwalker walk
# You can also try walkByTag to only export versions that are tagged.
gitwalker walkByTag
# Or use bare option to remove version name from exported folder names.
gitwalker walk --bare
gitwalker walkByTag --bare
```
5. Output folders are under the directory ```$HOME/.gitwalker/```