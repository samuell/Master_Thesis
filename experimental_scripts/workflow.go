package main

import (
	sp "github.com/scipipe/scipipe"
	spc "github.com/scipipe/scipipe/components"
)

func main() {
	wf := sp.NewWorkflow("DS_CPSign", 4)

	// Create target datasets
	createTargetDatasets := wf.NewProc("create_target_datasets", "python3 ../db_targets.py --database=../database --outdir={o:targets_dir}; echo 'done' > {o:doneflag}")
	createTargetDatasets.SetOut("targets_dir", "data/target_datasets")
	createTargetDatasets.SetOut("doneflag", "data/target_datasets.done")

	// Glob the target dataset files into a stream
	targetDirGlobber := spc.NewFileGlobberDependent(wf, "glob_target_datasets", "./data/target_datasets/[^.]*.tsv")
	targetDirGlobber.InDependency().From(createTargetDatasets.Out("doneflag"))

	// Balance the target datasets
	balanceTargetDatasets := wf.NewProc("balance_dataset", "python3 ../balancing_targets.py --dbfile=../database --infile={i:unbalanced} --outfile={o:balanced}; echo 'done' > {o:doneflag}")
	balanceTargetDatasets.In("unbalanced").From(targetDirGlobber.Out())
	balanceTargetDatasets.SetOut("balanced", `data/balanced_datasets/{i:unbalanced|basename|%.tsv}.balanced.tsv`)
	balanceTargetDatasets.SetOut("doneflag", "{o:balanced}.done")

	wf.Run()
}
