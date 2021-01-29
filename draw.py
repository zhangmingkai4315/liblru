import matplotlib.pyplot as plt

ns = [2**x for x in range(4, 17)];
get_not_in_lru = {
    'concurrentlru': [86.0,82.9,83.5,108,117,122,97.2,114,118,103,125,135,144],
    'slicelru': [27.6,24.1,35.9,64.1,67.6,70.8,50.5,79.9,90.7,76.0,103,133,145],
    'linklist': [38.7,44.8,43.0,70.3,77.5,82.6,60.7,87.8,97.1,74.5,96.6,112,116],
}

get_in_lru = {
    'concurrentlru': [65.3,64.5,69.6,76.1,95.4,104,107,112,115,121,122,129,132],
    'slicelru': [36.7,41.2,45.2,60.7,88.8,113,148,226,434,971,2210,5275,19257],
    'linklist': [26.7,28.0,28.2,36.6,57.7,69.4,71.8,73.9,78.0,81.0,84.7,88.3,93.8],
}
set_lru = {
    'concurrentlru-Set': [241,244,247,249,250,255,267,267,277,307,320,322,330],
    'slicelru-Set': [225,214,219,217,219,220,222,235,244,261,276,301,344],
    'linklist-Set': [212,212,212,210,212,216,219,226,235,249,259,262,280],
}
ax = plt.axes()
ax.set(xlabel='lru(size)', ylabel='time(ns)',
       title='GetNotInLRU');
for (label, values) in get_not_in_lru.items():
    plt.plot(ns, values, label=label)
plt.legend()
plt.show()

ax.set(xlabel='lru(size)', ylabel='time(ns)',
       title='GetInLRU');
for (label, values) in get_in_lru.items():
    plt.plot(ns, values, label=label)
plt.legend()
plt.show()

ax.set(xlabel='lru(size)', ylabel='time(ns)',
       title='SetLRU');
for (label, values) in set_lru.items():
    plt.plot(ns, values, label=label)
plt.legend()
plt.show()
