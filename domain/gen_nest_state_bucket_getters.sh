#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset
#$set -o xtrace

out='nest_state_bucket_getters.go'

echo "package domain

import (
	\"encoding/json\"
	\"fmt\"
	\"strings\"
)
/*
    THIS IS AN AUTOMATICALLY GENERATED FILE

    DO NOT EDIT!
*/
" > "${out}"

types="$( \
    grep 'const KeyPrefix.* = "[^"]\+"$' bucket_*.go \
    | sed 's/.*const KeyPrefix\([^ ]\+\).*/\1/' \
    | sort \
)"

for t in $types ; do
    tPl="${t}s"
    lc="$(tr '[:upper:]' '[:lower:]' <<< "${t:0:1}")""${t:1}"
    lcPl="${lc}s"
    if [[ "${tPl}" =~ ss$ ]] ; then
        tPl="${t}"
        lcPl="${lc}"
    fi
    tee -a "${out}" << EOF
// ${tPl} returns updated buckets items matching key
// prefix=KeyPrefix${t}.
func (ns *NestState) ${tPl}() ([]*${t}, error) {
	${lcPl}Slice := []*${t}{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefix${t}+".") {
			${lc} := &${t}{}
			if err := json.Unmarshal(v, &${lc}); err != nil {
				return nil, fmt.Errorf("unmarshalling to ${t}: %v", err)
			}
			${lcPl}Slice = append(${lcPl}Slice, ${lc})
		}
	}
	return ${lcPl}Slice, nil
}

EOF
done

gofmt -w "${out}"
