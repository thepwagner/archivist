# Archivist

Archivist is a quick hack job to track a collection of files.

* Files contents are indexed by SHA512 **and** BLAKE2b for integrity and deduplication.
* Index format is defined by protobuf, serialized as JSON for human interrogation.
* Index operates as a Git repository that is commit after updates.
* This makes some operations painfully slow compared to a "real" database!
* This makes development very forgiving (revert the index).

The index is split by "filesystem", which can be a mix of hot/cold storage devices.
The command line interface focuses on querying filesystems for a resource and identifying duplicates across filesystems.
