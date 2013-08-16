/* echo.h */

/********************************************************************************** 
 * Header file for the 64 bits optimized ansi C implementation of the ECHO hash 
 * function proposal.
 * Author(s) : Olivier Billet - Orange Labs - October 2008.
 **********************************************************************************/


#include <string.h>

typedef unsigned char BitSequence;
typedef unsigned long long DataLength;
typedef enum { 
  SUCCESS=0
 ,FAIL=1
 ,BAD_HASHBITLEN=2
 ,UPDATE_WBITS_TWICE=4
} HashReturn;


typedef struct {
  /* the state of the compression function
     it holds the chaining variable and the message */
  unsigned long long state[16*2];
  unsigned long long CV[8*2];
  /* the counter used by the domain extension */
  unsigned long long counter;
  /* the number of message bits currently stored in
  ** the message array */
  int message_bitlen;
  /* the size of the output hash in bits */
  int hashbitlen;
  /* the size of the chaining variable */
  int cv_blocks;
#ifdef SALT_OPTION
  unsigned char SALT[128];
#endif
} hashState;


#ifdef SALT_OPTION
HashReturn SetSalt(hashState* state, const BitSequence* SALT);
#endif

HashReturn Init(hashState *state, int hashbitlen);

HashReturn Update(hashState *state,
                  const BitSequence *data, DataLength databitlen);

HashReturn Final(hashState *state, BitSequence *hashval);

HashReturn Hash(int hashbitlen, const BitSequence *data,
                DataLength databitlen, BitSequence *hashval);
