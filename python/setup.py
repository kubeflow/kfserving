from setuptools import setup, find_packages
import os

setup(
    name='kfserving',
    version='0.1.0',
    author_email='ellisbigelow@google.com',
    license='../../LICENSE.txt',
    url='https://github.com/kubeflow/kfserving/model-servers/kfserver',
    description='Model Server for arbitrary python ML frameworks.',
    long_description=open('README.md').read(),
    python_requires='>3.4',
    packages=['kfserving.kfserver', 'kfserving.xgboost'],
    install_requires=[
        "tornado >= 1.4.1",
        "xgboost == 0.82",
        "argparse >= 1.4.0"
    ],
)
