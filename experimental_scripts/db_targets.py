import argparse
import pandas as pd
import os

#"--input", type = file, help = "input filename"
parser = argparse.ArgumentParser(description = 'Process the Chembl database')
parser.add_argument('targetdata', help = 'The database to take the targets from')
parser.add_argument('outfolder', help = 'The folder to produce the target files')
args = parser.parse_args()

with open(args.targetdata, 'r') as targetdata_file:
    df = pd.read_csv(targetdata_file, delimiter = '\t', lineterminator = '\n', header = 0)
    D={}

    for i in range(0, df.shape[0]):
        df.iloc[i, 0] = df.iloc[i, 0].replace('/','-')
        df.iloc[i, 0] = df.iloc[i, 0].replace(' ','_')
        df.iloc[i, 0] = df.iloc[i, 0].replace('(','_')
        df.iloc[i, 0] = df.iloc[i, 0].replace(')','_')
        df.iloc[i, 0] = df.iloc[i, 0].replace("'",'')

    for i in range(0, df.shape[0]):
        if df.iloc[i, 0] not in D:
            D[df.iloc[i, 0]] = []

    for j in range(0, df.shape[0]):
        D.setdefault(df.iloc[j,0],[]).append(df.iloc[j,1]+'\t')
        if df.iloc[j,2]=='IC50' and df.iloc[j,3]>10000:
            D.setdefault(df.iloc[j,0],[]).append('0\n')
        elif df.iloc[j,2]=='IC50' and df.iloc[j,3]<=10000:
            D.setdefault(df.iloc[j,0],[]).append('1\n')
        elif df.iloc[j,2]=='Ki' and df.iloc[j,3]>5000:
            D.setdefault(df.iloc[j,0],[]).append('0\n')
        elif df.iloc[j,2]=='Ki' and df.iloc[j,3]<=5000:
            D.setdefault(df.iloc[j,0],[]).append('1\n')
        elif df.iloc[j,2]=='Kd' and df.iloc[j,3]>5000:
            D.setdefault(df.iloc[j,0],[]).append('0\n')
        elif df.iloc[j,2]=='Kd' and df.iloc[j,3]<=5000:
            D.setdefault(df.iloc[j,0],[]).append('1\n')
        elif df.iloc[j,2]=='EC50' and df.iloc[j,3]>10000:
            D.setdefault(df.iloc[j,0],[]).append('0\n')
        elif df.iloc[j,2]=='EC50' and df.iloc[j,3]<=10000:
            D.setdefault(df.iloc[j,0],[]).append('1\n')

    if not os.path.isdir(args.outfolder):
        os.makedirs(args.outfolder)
    for key,value in D.items():
        outfile_path = args.outfolder + '/' + key + '.json'
        with open(outfile_path, 'w+') as outfile:
            print('Writing things to the outfile ' + outfile_path + ' ...')
            outfile.write('SMILES' + '\t' + 'FLAG' + '\n')
            outfile.write(''.join(value))
            outfile.close()

    #os.mkdir('only_non_binders')
