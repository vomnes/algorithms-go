import pandas as pd
import matplotlib.pyplot as plt
import numpy as np

data = pd.read_csv('sort_benchmark.csv', sep="\t", names=['name', 'iteration', 'runtime'])

_data = {
    'iteration': [],
}

types = []

for index, row in data.iterrows():
    if row["iteration"] not in _data["iteration"]:
        _data["iteration"].append(row["iteration"])
    if _data.get(row["name"]) == None :
        _data[row["name"]] = []
        types.append(row["name"])
    _data[row["name"]].append(row["runtime"])

df = pd.DataFrame(_data)

for type in types:
    plt.plot('iteration', type, data = df)
plt.xlabel('Iteration')
plt.ylabel('Runtime (sec.)')
plt.legend()
plt.show()
