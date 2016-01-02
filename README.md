# Handwritten image recognition

This neural network recognizes 28x28 handwritten images using a network of 824
neurons.


```sh
./digit say -in path/to/image-28x28.data -w path/to/weights.json -b path/to/biases.json
```

Images should be the matrix of intensities (0-255).
[https://github.com/wdalmut/mnist-handwritten-digits](https://github.com/wdalmut/mnist-handwritten-digits)

The network's accuracy is about 96%

