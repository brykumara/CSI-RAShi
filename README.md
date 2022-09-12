# CSIDH Lattice Reduction

As part of my MSc at UCL, I implemented lattice reduction algorithm for CSIDH, an isogeny-based post-quantum cryptography standard. This is a key functionality that allows users to start with an integer as the secret and convert it into an exponent vector used in CSIDH as the secret key. Crucially, we need this feature to conduct digital signatures and multiparty computations with the other alternative being Fiat-Shamir aborts. 

This implementation is based on the CSI-FiSh digital signature (https://eprint.iacr.org/2019/498) and the original aim of the thesis was to implement a distributed key generation based on CSIDH. However, upon learning that random integers tino exponent vectors were not supported on the CIRCL library on Go for CSIDH, the research shifted to examining this key feature. 

In this implementation, we have worked with the CIRCL library to implement lattice reduction using the BKZ-40, BKZ-50, and HKZ bases alongside the L1 and L2 norms. The results indicate that this procedure is extremely inefficient, being on par with other CSIDH implementations, and is a bitter pill that needs to be swallowed only once per party for multiparty computations. Without this feature available, users would have to resort to Fiat-Shamir aborts which also has a fairly long run time. 

This dissertation was supervised by Prof. Philipp Jovanovic and Ms. Maria Real-Corte Santos over the summer of 2022.
