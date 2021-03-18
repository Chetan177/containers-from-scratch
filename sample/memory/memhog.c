/*
 * Credit:
 *  https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/6/html/resource_management_guide/sec-memory
 *  Original source code provided by Red Hat Engineer František Hrbata.
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#define KB (1024)
#define MB (1024 * KB)
#define GB (1024 * MB)

#define LIMIT (5.0 * MB)

int main(int argc, char *argv[])
{
	char *p;
	double alloc = 0;

again:
	while ((alloc+GB) < LIMIT && (p = (char *)malloc(GB))) {
		alloc += GB;
		memset(p, 0, GB);
	}

	while ((alloc+MB) < LIMIT && (p = (char *)malloc(MB))) {
		alloc += MB;
		memset(p, 0, MB);
	}

	while ((alloc+KB) < LIMIT && (p = (char *)malloc(KB))) {
		alloc += KB;
		memset(p, 0, KB);
	}

	sleep(1);

	goto again;

	return 0;
}
