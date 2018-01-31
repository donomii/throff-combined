package throfflib
func BootStrapString( ) string {




	var str string
	str =`

	
	
	TESTBLOCK [ PRINTLN [ Bootstrap complete, ready for commands ] ]

	
	
	DEFINE SUBSHELL => LAMBDA [
	IF EQUAL OS windows [ 
		STARTPROCESS C:\Windows\system32\cmd.exe args 
		PRINTLN args
		BIND args => UNSHIFTARRAY /c CMD 
	] [
		STARTPROCESS /bin/sh args 
		PRINTLN args
		BIND args => UNSHIFTARRAY /c CMD 
	]
	PRINTLN CMD 
	BIND CMD => 
]
	
DEFINE IFFY => [ IF TRUTHY ]

DEFINE TRUTHY => [

       CASE A[
               EQUAL 0 X             ... FALSE
               LESSTHAN X 0          ... FALSE
               EQUAL FALSE X         ... FALSE
               NOT EQUAL 0  LENGTH X ... TRUE
               DEFAULT               ... TRUE
       ]A

       ARG X =>
]

DEFINE THROW => [
    ACTIVATE/CC ERRORHANDLER A[ TRUE  MSG ]A
    ARG MSG =>
]

DEFINE CATCH => [
    IF HANDLER [
        CALL ERRLAMBDA MSG
    ] [ ]

    BIND MSG => CAR CDR RET
    BIND HANDLER => CAR RET
    BIND RET =>
    CALL/CC [
        INSTALLDYNA CC ->FUNC  [ A[ FALSE AAAA AAAA  ]A CALL ALAMBDA ]
        ARG CC =>
    ]
    ARG ALAMBDA =>
    ARG ERRLAMBDA =>
]


BIND DEFAULT => TRUE

DEFINE CASE => [
  WHEN NOT EMPTY? ARR [
  IF CALL TEST [ CALL FUNC ]
  [ CASE CDR CDR ARR ]
  ]

  BIND FUNC => CAR CDR ARR
  BIND TEST => CAR ARR
  ARG ARR =>
]

DEFINE ... => [ ]
DEFINE EMPTY? => [ EQUAL 0 LENGTH ]



DEFINE STRING-JOIN => [
	FOLD [ STRING-CONCATENATE STRING-CONCATENATE SWAP A SWAP ] CAR ARR CDR ARR
	ARG ARR =>
	ARG A =>
]

DEFINE CAR => [ DROP SWAP SHIFTARRAY   ]

DEFINE CDR => [ DROP SHIFTARRAY ]


DEFINE STRING-CONCATENATE* => [
	FOLD [ STRING-CONCATENATE ] [ ]
]

DEFINE APPEND => [
	TARGET
	ITERATE THIN [ REBIND TARGET => ARRAYPUSH TARGET ] SOURCE
	ARG SOURCE =>
	ARG TARGET =>
]
DEFINE PL => [
	IF NOT EQUAL 1 X
		[ STRING-CONCATENATE WORD s TOK ]
		WORD
	BIND WORD =>
	BIND X =>
]

DEFINE REVERSE => [
	NEW
	ITERATE THIN [ REBIND NEW => UNSHIFTARRAY SWAP NEW ] OLD
	BIND NEW => NEWARRAY
	ARG OLD =>
]



	 DEFINE AND => [ IF EQUAL A B [ IF EQUAL A TRUE TRUE FALSE ] FALSE : B : A ]

DEFINE PROMISE => [
        ->FUNC [
                VAL
                WHEN NOT DONE
                        THIN [
				REBIND F => FALSE
                                REBIND VAL => F
                                REBIND DONE => TRUE
                        ]
        ]
	ARG F =>
]

DEFINE CALLA => [
    CALL [
        ->FUNC THIN [
                VAL
                WHEN NOT EQUAL DONE 1
                        THIN [
                                REBIND VAL => READQ RQ
                                REBIND DONE => 1
                                COMMENT [ DUMP RQ
                                PRINTLN A[ !!!!!! [ CALLA READING RETURN QUEUE ] URL !!!!!!! ]A ]
                        ]
        ]
        REBIND RQ => RQ
        REBIND DONE => DONE
        REBIND VAL => VAL
    ]
    WRITEQ AC H[ FARG => URL RET => RQ ]H
    COMMENT [ DUMP AC
    PRINTLN A[ [ CALLA WRITING TO QUEUE ] URL ]A ]
    BIND RQ => NEWQUEUE
    BIND DONE => 0
    BIND VAL => FALSE
    BIND URL =>
    BIND AC =>
]

DEFINE ACTORREC => [
                    COMMENT [ PRINTLN [ CALLED RECURSE in ACTORREC ] ]
                    ACTORREC
                    COMMENT [ PRINTLN [ CALLING RECURSE in ACTORREC ] ]
                    USERFUNC WQ
                    WRITEQ RETQ CALL USERFUNC  FARG STATE
                    COMMENT [ DUMP RETQ
                    PRINTLN [ ACTOR WRITING TO QUEUE ] ]
                    BIND RETQ => GETHASH RET => PARAMS
                    BIND FARG => GETHASH FARG => PARAMS
                    BIND PARAMS => READQ WQ
                    COMMENT [ DUMP WQ
                    PRINTLN [ ACTOR WAITING ON QUEUE ] ]
                    BIND STATE =>
                    BIND WQ =>
                    BIND USERFUNC =>
                    COMMENT [ PRINTLN [ Recursing ACTORREC ] ]
]

DEFINE ACTOR => [
        WQ
        THREAD [
                ACTORREC USERFUNC WQ STARTPROCESS /bin/echo  [ . ]
        ]
        BIND WQ NEWQUEUE
        BIND USERFUNC =>
]


DEFINE MIME => [
        ->FUNC [ WRITEQ Q ]
        THREAD [
                FOREVER [ CALL USERFUNC  READQ Q ]
        ]
        BIND Q NEWQUEUE
        BIND USERFUNC =>
]

DEFINE FOREVER => [ FOREVER CALL DUP ]


TESTBLOCK [

TEST  wonkyA   1  [ WITH STATEMENT ]

 WITH   A[ wonkyA ]A FROM TESTHASH

BIND TESTHASH => H[ wonkyA 1 B 2 C 3 ]H

]

DEFINE WITH =>   MACRO [

	ITERATE THIN [


		BIND  with_N GETHASH with_N  with_AHASH

		REBIND with_N =>


	]  with_NAMES

	d binding > with_NAMES <  from hash > with_AHASH < ;

	BIND with_N => WHOOPS
	ARG with_AHASH =>
	ARG with_FROM =>
	ARG with_NAMES => ->ARRAY



]






DEFINE -f => [ RUNSTRING  SLURPFILE  SWAP  ENVIRONMENTOF HERE TOK ]
DEFINE -s => [ SAFETYON ]

DEFINE SLURPFILE => [ FOLD [ STRING-CONCATENATE SWAP STRING-CONCATENATE CRLF ] ->STRING [ ] MAPFILE [  ] ]

DROP [



SEARCHBYTES ->BYTES [ HTTP ] map
ITERATE [ EMIT GETBYTE SWAP map ] RANGE 0 255

: map => MMAPFILE apache.log



]


DEFINE DUMPBYTES => [

	ITERATE [ EMIT GETBYTE SWAP B ] RANGE 0 ADD -1 LENGTH B

	: B =>
]

DEFINE COMPAREBYTES => [
PRINTLN DUP
	returns SUCCESS
	ITERATE THIN [


 		STATEMENT [ COMPAREBYTES INTERNAL ] [ IF EQUAL
				GETBYTE I SMALLBYTES
				GETBYTE ADD I INDEX LARGEBYTES
				THIN [ PRINTLN [ MATCH SUCCESS ]  REBIND SUCCESS => TRUE ]
				THIN [ PRINTLN [ MATCH FAIL ] REBIND SUCCESS => FALSE ]
		]

		p Comparing bytes at I and ADD I INDEX ;
		REBIND I =>
	]

	BIND I => -1
	RANGE 0 ADD -1 LENGTH SMALLBYTES

	BIND SUCCESS => WHOOPS
	ARG LARGEBYTES =>
	ARG INDEX =>
	ARG SMALLBYTES =>
]

DEFINE SEARCHBYTES => [

	FOR 0
	ADD -1 SUB LENGTH LARGEBYTES LENGTH SMALLBYTES
	[
		IF   COMPAREBYTES SMALLBYTES I LARGEBYTES [
			p found search target at I ;
		] [ ]


		ARG I =>
	]


	p length a: LENGTH SMALLBYTES  b: LENGTH LARGEBYTES, iterating to: ADD  SUB LENGTH LARGEBYTES LENGTH SMALLBYTES  ;
	ARG LARGEBYTES =>
	ARG SMALLBYTES =>
]


DEFINE MAPFILE => [
	returns RESULT

	CLOSEFILE FHANDLE
	BIND RESULT => MAPFILERECURSE
					NEWARRAY
					GETFUNCTION MFUNC TOK
					READER
	BIND FHANDLE => BIND READER => OPENFILE FILENAME
	ARG FILENAME =>
	ARG MFUNC => UNFUNC
]


DEFINE MAPFILERECURSE =>  [
	IF NOT EQUAL  ->STRING LINE  ->STRING FALSE TOK   [
		MAPFILERECURSE
			ARR
			GETFUNCTION MFUNC TOK
			FHANDLE
		REBIND ARR => ARRAYPUSH ARR CALL MFUNC LINE
	]
	[ returns  ARR  ]



	BIND LINE => READFILELINE FHANDLE




	ARG FHANDLE =>
	ARG MFUNC => UNFUNC
	ARG ARR =>
]

: SW => [
	PRINTLN [ ------------------------------- ]
	PRINTLN DUP
	PRINTLN [ ------------------------------- ]
]

TESTBLOCK [

	TEST  MERGESORT [ LESSTHAN ] A[ 6 3 5 2 1 4 ]A  A[ 1 2 3 4 5 6 ]A ->STRING [ MERGESORT ]

]

DEFINE MERGESORT => [


	IF LESSTHAN  LENGTH ARR 2 [
		returns ARR
	]
	[
		returns MERGE GETFUNCTION MFUNC TOK
					MERGESORT GETFUNCTION MFUNC TOK  LEFT
					MERGESORT GETFUNCTION MFUNC TOK  RIGHT
					A[ ]A
		: LEFT => 	SLICE 0 SUB HALFWAY 1 ARR
		: RIGHT =>	SLICE
						HALFWAY
						SUB LENGTH ARR 1
						ARR
		: HALFWAY => FLOOR DIVIDE LENGTH ARR 2

	]

	d Merge sorting ARR ;

	: ARR =>
	: MFUNC =>
]


TESTBLOCK [


	TEST SLICE-STRING 1 3 ABCDEFG BCD    ->STRING [ SLICE-STRING ]

]

DEFINE SLICE-STRING => [

	returns NEWSTR

	ITERATE THIN [

		REBIND NEWSTR => STRING-CONCATENATE NEWSTR  GETSTRING I STR
		REBIND I =>

	] RANGE START END

	: I => 0
	: NEWSTR => ->STRING [ ]


	d Slicing START to END in STR ;


	: STR =>
	: END =>
	: START =>

]


TESTBLOCK [
	TEST SLICE 1 3 A[ The quick brown fox jumped over the lazy dog ]A  A[ quick brown fox ]A    ->STRING [ SLICE ]

]


DEFINE SLICE => [

	returns NEWARR

	ITERATE THIN THIN [

		REBIND NEWARR => ARRAYPUSH NEWARR GETARRAY I ARR
		REBIND I =>

	] RANGE START END

	: I => 0
	: NEWARR => NEWARRAY


	d Slicing START to END in ARR ;


	: ARR =>
	: END =>
	: START =>

]

TESTBLOCK [

	TEST  MERGE [ LESSTHAN ] A[ 1 3 5 ]A A[ 2 4 6 ]A A[ ]A  A[ 1 2 3 4 5 6 ]A ->STRING [ MERGE ]
]


DEFINE MERGE => [


	IF LESSTHAN 0 LENGTH L1 THIN [
		IF LESSTHAN 0 LENGTH L2 THIN [
			MERGE GETFUNCTION MFUNC TOK L1 L2 ACCUM
			IF CALL MFUNC A B THIN [
				REBIND ACCUM => ARRAYPUSH ACCUM A
				REBIND L1 => DROP SHIFTARRAY L1
			]
			THIN [

				REBIND ACCUM => ARRAYPUSH ACCUM B
				REBIND L2 => DROP SHIFTARRAY L2

			]

			: A => GETARRAY 0 L1
			: B => GETARRAY 0 L2
		]

		THIN [

			MERGE GETFUNCTION MFUNC TOK L1 L2 ACCUM
			REBIND ACCUM => ARRAYPUSH ACCUM A
			REBIND L1 => DROP SHIFTARRAY L1
			: A => GETARRAY 0 L1
		]
	]

	THIN [

		IF LESSTHAN 0 LENGTH L2 THIN [
				MERGE GETFUNCTION MFUNC TOK L1 L2 ACCUM
				REBIND ACCUM => ARRAYPUSH ACCUM B
				REBIND L2 => DROP SHIFTARRAY L2
				: B => GETARRAY 0 L2
			]
			THIN [
				d Merge returning ACCUM ;
				ACCUM
			]


	]


	d ----Merging-------- ;
	d l1 L1  length LENGTH L1 ;
	d l2 L2  length LENGTH L2 ;
	d accum ACCUM ;
	d ------------------- ;


	: ACCUM =>
	: L2 =>
	: L1 =>
	: MFUNC =>


]

TESTBLOCK [
	TEST MULTIBUBBLESORT A[ A[ 3 8 2 20 5 ]A A[ AA BB CC DD EE ]A  ]A	A[ A[ 2   3   5   8   20 ]A A[ CC   AA   EE   BB   DD ]A ]A ->STRING [ MULTIBUBBLESORT ]
]

DEFINE MULTIBUBBLESORT => [

	IF CHANGED [
		MULTIBUBBLESORT NEWARR
		]
		[
			returns NEWARR
		]
	: NEWARR => : CHANGED => MULTIBUBBLESTEP ARR



	: ARR =>
]

TESTBLOCK [

TEST  DROP MULTIBUBBLESTEP DROP MULTIBUBBLESTEP A[ A[ 3 8 2 20 5 ]A A[ AA BB CC DD EE ]A  ]A	A[ A[ 2   3   5   8   20 ]A A[ CC   AA   EE   BB   DD ]A ]A  [ MULTIBUBBLESTEP ]
]

DEFINE MULTIBUBBLESTEP => [

	returns  CHANGED ARRS


	ITERATE THIN [
		IF LESSTHAN
			GETARRAY ADD 1 I ARR
			GETARRAY I ARR
			THIN [
				STATEMENT [ MULTIBSTEP REBIND 2 ] [ REBIND CHANGED => TRUE ]
				STATEMENT [ MULTIBSTEP REBIND 1 ] [ REBIND ARRS => MAP [ SWAPARR ADD 1 I I  ] ARRS ]
			]
			[ ]

		REBIND I =>
	]  RANGE 0 SUB LENGTH ARR 2

	: I => 0
	: CHANGED => FALSE
	: ARR => GETARRAY 0 ARRS

	: ARRS =>
]

TESTBLOCK [

	TEST BUBBLESORT A[ 3 8 2 20 5 ]A 	A[ 2   3   5   8   20 ]A ->STRING [ BUBBLESORT ]

]

DEFINE BUBBLESORT => [

	IF CHANGED [
		BUBBLESORT NEWARR
		]
		[
			returns NEWARR
		]


	ARG NEWARR =>
	ARG CHANGED => BUBBLESTEP ARR



	ARG ARR =>
]

TESTBLOCK [
	TEST DROP BUBBLESTEP DROP BUBBLESTEP A[ 3 8 2 20 5 ]A 	A[ 2   3   5   8   20 ]A ->STRING [ BUBBLESTEP ]
]

DEFINE BUBBLESTEP => [

	returns CHANGED ARR

	ITERATE THIN [
		IF LESSTHAN
			GETARRAY ADD 1 I ARR
			GETARRAY I ARR THIN [
				REBIND CHANGED => TRUE
				REBIND ARR => SWAPARR ADD 1 I I ARR
			]
			[ ]

		REBIND I =>
	]  RANGE 0 SUB LENGTH ARR 2

	: CHANGED => FALSE
	: I => UNINITIALISED
	: ARR =>
]

TESTBLOCK [
	TEST RANGE 1 3 A[ 1 2 3 ]A ->STRING [ RANGE ]
	]

DEFINE RANGE => [

	returns ARR
	COUNTDOWN2ZERO
	COUNT THIN [
		REBIND ARR => UNSHIFTARRAY  ADD MIN I ARR
		REBIND I TOK
	]


	: I => 0
	: ARR => NEWARRAY

	: COUNT => SUB MAX MIN
	: CHANGED => FALSE

	: MAX =>
	: MIN =>
]



DEFINE INVERTHASH => [
	ITERATE2 [
		SETHASH VAL KEY NEWHASH
		: KEY => : VAL =>
	]
]


TESTBLOCK [

	TEST GETARRAY 0 swappedSpirits ->STRING [ vodka ] ->STRING [ SWAP ]

	TEST GETARRAY 1 swappedSpirits ->STRING [ rum ] ->STRING [ SWAP ]



	: swappedSpirits => SWAPARR 0 1 moreSpirits



	: moreSpirits => UNSHIFTARRAY ->STRING [ rum ] spirits

	TEST GETARRAY 0 spirits  [ vodka ]  [ ARRAYPUSH and GETARRAY ]

	: spirits => ARRAYPUSH NEWARRAY ->STRING [ vodka ]


]

COMMENT [ Swaps two elements in an array... and does two full copies of the array! ]
DEFINE SWAPARR =>  [
	returns ARR2

	: ARR2 => SETARRAY B TEMPVAL ARR1

	: ARR1 => SETARRAY A GETARRAY B ARR ARR

	: TEMPVAL => GETARRAY A ARR
	: ARR =>
	: B =>
	: A =>
]

: SORT [


	:ARR =>
]



: VALS => [ DROP KEYS/VALS ]
: KEYS => [ DROP SWAP KEYS/VALS ]
: KEYS/VALS => [  ARR2 ARR ITERATE2 THIN [ REBIND ARR => ARRAYPUSH ARR  REBIND ARR2 => ARRAYPUSH ARR2  ] KEYVALS : ARR => A[ ]A : ARR2 => A[ ]A  ]




->STRING [ Stack empty. ]


DEFINE HELP => [ ITERATE [ PRINTLN ] A[
.
[ Welcome to THROFF ]
.
[ Throff is a minimal scripting language and an ultra-lightweight interpreter. Throff provides powerful and convenient scripting access from the command line. ]
.
.
[ SHELLHELP   - Using the command line ]
[ QUICKSTART  - How to Throff, with examples ]
[ BUILTINS    - Arrays, hashes, if then else, while, map, filter, fold etc ]
.
.
.
]A
: . => A[ ]A
: TAB => STRING-CONCATENATE STRING-CONCATENATE SPACE SPACE STRING-CONCATENATE SPACE SPACE
]

DEFINE SHELLHELP => 	[ ITERATE [ EMIT SPACE EMIT ] A[
.
[ Type your program on the command line, and the result will be printed above the ready prompt. ]
Ready> ADD 2 3
5
Ready> 

[ Throff is a stack based language.  The top item of the stack is always printed above the Ready> prompt.  All functions work on the stack, so you can type your command in parts, and examine it. ]

Ready> 2
2
Ready> 3
3
Ready> ADD
5
Ready>

]A
: . => A[ ]A
: TAB => STRING-CONCATENATE STRING-CONCATENATE SPACE SPACE STRING-CONCATENATE SPACE SPACE
]

DEFINE  QUICKSTART => 	[ ITERATE [ EMIT SPACE EMIT ] A[

. Throff is a catenative language, so you don't need brackets, braces or commas to do function calls.  This makes it easy and quick to type commands, especially in the command shell.
.
. TAB Ready> [ p ADD 1 1 ; ]
.
. TAB 2
.
. Functions usually have a fixed number of arguments, and Throff keeps track of them.
.
. TAB Ready> [ p  SUB 1 MULT 2 ADD 1 1 ; ]
.
. TAB -3
.
. Throff commands run from right to left.  The result of the previous command is the arguments to the next
.
. Throff has a very minimal parser.  Control characters must never touch anything else
.
.  Good: TAB [ PRINTLN [ Hello ] ]
.
.  Bad: TAB [PRINTLN [Hello]]
.
.
]A
: . => STRING-CONCATENATE CR NEWLINE
: TAB => STRING-CONCATENATE STRING-CONCATENATE SPACE SPACE STRING-CONCATENATE SPACE SPACE
]






DEFINE test => [
		TEST TESTVAR EXPECTED DESCRIPTION
		ARG TESTVAR TOK
		ARG EXPECTED TOK
		ARG DESCRIPTION => STRING
]









DEFINE ADD1 => [ ADD  1 ]
DEFINE SUB1 => [ SUB SWAP 1  ]
DEFINE FI => [ ]
DEFINE ELSE => [ ]

TESTBLOCK [

	TEST GETARRAY 0 spirits ->STRING [ vodka ] ->STRING [ PUSH something ONTO ]

	REBIND spirits => PUSH ->STRING [ vodka ] ONTO spirits

	: spirits => NEWARRAY

]


DEFINE PUSH => GETFUNCTION PUSHARRAY TOK
DEFINE PUSHARRAY => [ ARRAYPUSH  SWAP  ]
DEFINE ONTO => [  ]
DEFINE GET => [ DROP OF ]

DEFINE OF => [

	IF ( NOT EQUAL ->STRING HASH TOK  GETTYPE A ) [
			OF TOK
			GETHASH A  AHASH
			ARG AHASH  =>
			ARG of =>
		]
		[
			OF TOK A
		]

	ARG A =>
	TROFF
]

TESTBLOCK [
	TEST KEY pages 	[ HASHITERATE key  ]
	TEST VAL 132 	[ HASHITERATE value ]

	HASHITERATE  bookdata THIN [  REBIND VAL => REBIND KEY => ]

	: KEY => FAIL
	: VAL => FAIL


	: bookdata => SETHASH pages 132 NEWHASH

]


DEFINE  HASHITERATE => [
	ITERATE2 THIN [
		CALL  AFUNC  AKEY  AVAL
		ARG AVAL =>
		ARG AKEY =>

	]



	KEYVALS AHASH

	d keyvals KEYVALS AHASH ;
	ARG AFUNC =>  UNFUNC
	ARG AHASH =>
]




TESTBLOCK [


TEST ACCUM 3 ->STRING [ ITERATE2 ]

ITERATE2 THIN [ REBIND ACCUM => ADD 1 ACCUM DROP DROP ] A[ 1 2 3 4 5 6 ]A

: ACCUM => 0

]


TESTBLOCK [

TEST ACCUM 		ITERATE THIN [ REBIND ACCUM => SUB ACCUM 1 DROP ] A[ 1 2 3 4 5 ]A 		0 ->STRING [ ITERATE ]

: ACCUM => 5

]



DEFINE ITERATE2RECURSE => [

	IF LESSTHAN I LENGTH ARRAY1 [
		ITERATE2RECURSE ADD 2 I   ARRAY1  MFUNC

		STATEMENT [ ITERATE2RECURSE internal function call ] [    CALL  MFUNC  ELEM1   ELEM2    ]

		d args are GETFUNCTION ELEM1 TOK  and GETFUNCTION ELEM2 TOK ;

		: ELEM2 =>  GETARRAY ADD 1 I ARRAY1
		: ELEM1 =>  GETARRAY I ARRAY1

	]
	[ ]

	d function is  MFUNC  ;
	d .iterate2recurse. I th element of array of length LENGTH ARRAY1 ;

	ARG MFUNC => UNFUNC
	ARG ARRAY1 =>
	ARG I =>
]

DEFINE ITERATE2 => [
	ITERATE2RECURSE 0 SWAP
]


DEFINE ITERATESLICE => [

	IF LESSTHAN END START [
	]
	[

		ITERATESLICE MFUNC ADD 1 START END ARRAY1
		STATEMENT [ Internal iterate call ] [ CALL MFUNC  GETARRAY START ARRAY1 ]

	]

	ARG ARRAY1 TOK
	ARG END TOK
	ARG START TOK
	ARG MFUNC  TOK
]


DEFINE ITERATE => [
	ITERATESLICE MFUNC 0 ADD -1 LENGTH ARRAY1 ARRAY1

	ARG ARRAY1 TOK
	ARG MFUNC  TOK UNFUNC
]



DEFINE FORRECURSE => [

	IF LESSTHAN END START [ ]
	[

		FORRECURSE  ADD 1 START END  MFUNC
		STATEMENT [ Internal FOR call ] [ CALL MFUNC  START ]

	]

	ARG MFUNC  TOK
	ARG END TOK
	ARG START TOK

]




DEFINE FOR => [
	FORRECURSE START END MFUNC


	ARG MFUNC  TOK
	ARG END TOK
	ARG START TOK
]


TESTBLOCK [
	TEST   FOLD [ ADD ] 0   A[ 1 2 3 ]A    6   ->STRING [ FOLD ]
]

COMMENT [ FIXME OH GOD THIS IS BAD ]
DEFINE FOLD => [
	returns ACCUM
	COUNTDOWN2ZERO
	SUB LENGTH ARRAY1 1 THIN [
		REBIND ACCUM => CALL MFUNC
							GETARRAY SUB SUB LENGTH ARRAY1 I 1 ARRAY1
							ACCUM


		REBIND I TOK
	]


	ARG I => UNINITIALISED TOK
	ARG ARRAY1 =>
	ARG ACCUM =>
	ARG MFUNC  =>
]

TESTBLOCK [
	TEST   FILTER [ NOT EQUAL 2  ]  A[ 1 2 3 ]A   A[ 1 3 ]A        ->STRING [ FILTER ]
]

DEFINE FILTER => [
	returns OUTPUT
	COUNTDOWN2ZERO
		SUB LENGTH ARRAY1 1 THIN [
			IF ( CALL MFUNC ELEM )
				THIN [ REBIND OUTPUT TOK ARRAYPUSH ( OUTPUT  ELEM ) ]
				[ ]
			REBIND ELEM => GETARRAY SUB SUB LENGTH ARRAY1 I 1 ARRAY1
			REBIND I TOK
		]

	ARG I => 0
	ARG OUTPUT => NEWARRAY
	ARG ELEM => UNINITIALISED
	ARG ARRAY1 TOK
	ARG MFUNC  TOK
]


TESTBLOCK [
	TEST   MAP [ ADD 1  ]  A[ 1 2 3 ]A   A[ 2 3 4 ]A         [ MAP ]
	]

DEFINE MAP => [
	returns OUTPUT
	COUNTDOWN2ZERO
	SUB LENGTH ARRAY1 1 THIN [
		REBIND OUTPUT TOK ARRAYPUSH OUTPUT
			CALL MFUNC GETARRAY  SUB SUB LENGTH ARRAY1 I 1 ARRAY1
		REBIND I TOK

	]


	BIND I => 0
	ARG OUTPUT => NEWARRAY

	ARG ARRAY1 TOK
	ARG MFUNC  TOK UNFUNC
]

DEFINE STACKMAP => [
	RECURSE
	ARG RECURSE => [
		IF EQUAL ; PICK 2
		[ DROP ]
		[
			RECURSE
			FUNC
		]
	]

	ARG FUNC =>
]

COMMENT [ p Testing variable print command ... Test passed! ; ]

DEFINE ; => [ ; TOK ]


DEFINE d => [

		IF EQUAL  ; TOK X
		[   IF EQUAL 0 1 [
				PRINTLN [ ]
			]
			[ ]
		]
		[ 	IF EQUAL 1 1 [
				d
			]
			[ d EMIT SPACE EMIT X ]
		]

		ARG X => ->STRING
]


DEFINE pn => [
	IF EQUAL  ; TOK X
		[ PRINTLN [ ]  ]
		[ p EMIT SPACE EMIT X ]

		ARG X => ->STRING
]


DEFINE p => [
	IF EQUAL  ; TOK X
		[  ]
		[ p EMIT SPACE EMIT X ]

		ARG X => ->STRING
]

TESTBLOCK [
	TEST GETHASH G TESTHASH H ->STRING [ H[ ]H syntax ]
	TEST GETHASH e TESTHASH F ->STRING [ H[ ]H syntax ]


	: TESTHASH =>  H[  e F G H  ]H
]


DEFINE H[RECURSE => [
	IF EQUAL ->STRING ]H TOK ->STRING KEY
		[ AHASH ]
		[
			H[RECURSE AHASH
			REBIND AHASH => HASHSET AHASH KEY GETFUNCTION VALUE TOK
		]



	ARG VALUE =>
	ARG KEY =>
	ARG AHASH =>
]

DEFINE H[ => [
   H[RECURSE NEWHASH
]


DEFINE ]H => [ ->STRING ]H TOK BUFFERVAL TOK ]


TESTBLOCK [

	TEST KEYVALS bookdata A[ pages 132 ]A ->STRING [ KEYVALS ]


	: bookdata => SETHASH pages 132 NEWHASH


]
TESTBLOCK [
	TEST GETARRAY 1 grains rye ->STRING [ A[ ]A syntax ]

	: grains => A[ wheat rye barley ]A

]
DEFINE A[RECURSE => [

	IF EQUAL ->STRING ]A TOK ->STRING GETFUNCTION VAL TOK
		[ ARR  ]
		[
			A[RECURSE ARR
			REBIND ARR => ARRAYPUSH ARR GETFUNCTION VAL TOK
		]


	ARG VAL =>
	ARG ARR =>
]



DEFINE A[ => [
   A[RECURSE NEWARRAY
]

BIND ]A => ->STRING ]A TOK


DEFINE PRINTX => [
	PRINTLN [ ]
	REPEAT SWAP [ EMIT SPACE EMIT ]
]

COMMENT [ REPEAT 3 [ PRINTLN ->STRING [ Testing repeat 3 ] ] ]



DEFINE REPEAT => [
	COUNTDOWN TIMES
	[
		CALL MFUNC
		ARG I TOK
	]

	ARG MFUNC  TOK ->LAMBDA
	ARG TIMES  TOK
]

DEFINE COUNTDOWN => [ COUNTDOWN2ZERO SUB SWAP 1 ]

COMMENT [ COUNTDOWN2ZERO 5 [ PRINTLN EMIT SPACE EMIT ->STRING [ COUNT: ] ] ]

DEFINE COUNTDOWN2ZERO => [
	IF LESSTHAN ( COUNT  ZERO )
		[  ]
		[
			COUNTDOWN2ZERO ( SUB ( COUNT 1 )  GETFUNCTION ( FUNC TOK ) )
			CALL FUNC COUNT
		]

	ARG FUNC TOK  UNFUNC
	ARG COUNT TOK
]

DEFINE FUNCTION? => [
	IF EQUAL TYPE CODE TOK
		[ ]
		[ ERROR A[ LOCATIONOF TYPE [ EXPECTED FUNCTION BUT GOT SOMETHING ELSE INSTEAD AT:  ] ]A ]

	ARG TYPE TOK GETTYPE DUP
]

DEFINE ARRAY => [
	IF EQUAL TYPE ARRAY TOK
		[ ]
		[ ERROR [ EXPECTED ARRAY BUT GOT SOMETHING ELSE INSTEAD ] ]

	ARG TYPE TOK GETTYPE DUP
]

DEFINE STRING => [
	IF EQUAL TYPE STRING TOK
		[ ]
		[ ERROR [ EXPECTED STRING BUT GOT SOMETHING ELSE INSTEAD ] ]

	ARG TYPE TOK GETTYPE DUP
]

TESTBLOCK [
	TEST WHEN [ EQUAL 1 1 ] [ ->STRING [ PASS ] ]   ->STRING [ PASS ] ->STRING [ WHEN statement with function for conditional ]
	TEST WHEN EQUAL 1 1 	[  WHENPASS  ]   WHENPASS  ->STRING [ WHEN statement with function for true branch ]
	TEST WHEN EQUAL 1 1  	  ->STRING [ PASS ]     ->STRING [ PASS ] ->STRING [ WHEN statement with constant true branch ]
]

DEFINE WHEN => [
	IF CALL VAL
		[ CALL ACTION ]
		[ ]

		: ACTION TOK
		: VAL TOK
]

COMMENT [
    DEFINE ERROR => [ CALL/CC [ ERRORHANDLER ] PRINTLN ]

DEFINE ERRORHANDLER => [
	INTERPRET
	PRINTLN [ An error has occurred, dropping you to the debugger ]
]
]

TESTBLOCK [
	TEST OPEN-FUNCTION-CHAR 	NUM2CHAR 93 	->STRING [ OPEN-FUNCTION-CHAR ]
	TEST CLOSE-FUNCTION-CHAR 	NUM2CHAR 91 	->STRING [ CLOSE-FUNCTION-CHAR ]
]
BIND OPEN-FUNCTION-CHAR => NUM2CHAR 93
BIND CLOSE-SQUARE-BRACE-CHAR => NUM2CHAR 93
BIND CLOSE-FUNCTION-CHAR => NUM2CHAR 91
BIND OPEN-SQUARE-BRACE-CHAR => NUM2CHAR 91
BIND CLOSE-CURLY => ->STRING } TOK
BIND OPEN-CURLY => ->STRING { TOK
BIND , => ,

TESTBLOCK [

	TEST GETARRAY 1 moreSpirits ->STRING [ vodka ] ->STRING [ clone ARRAY ]
	TEST GETARRAY 0 moreSpirits ->STRING [ rum ] ->STRING [ clone ARRAY ]

	TEST GETARRAY 1 alteredSpirits ->STRING [ whiskey ] ->STRING [ SETARRAY ]

	: alteredSpirits => SETARRAY 1 ->STRING [ whiskey ] moreSpirits

	TEST GETARRAY 1 moreSpirits ->STRING [ vodka ] ->STRING [ clone ARRAY ]
	TEST GETARRAY 0 moreSpirits ->STRING [ rum ] ->STRING [ clone ARRAY ]


	TEST DROP SWAP POPARRAY moreSpirits ->STRING [ vodka ] ->STRING  [ POP ]

	TEST GETARRAY 1 moreSpirits ->STRING [ vodka ] ->STRING [ clone ARRAY ]
	TEST GETARRAY 0 moreSpirits ->STRING [ rum ] ->STRING [ clone ARRAY ]


	TEST DROP SWAP SHIFTARRAY moreSpirits ->STRING [ rum ] ->STRING [ SHIFTARRAY ]

	TEST GETARRAY 1 moreSpirits ->STRING [ vodka ] ->STRING [ UNSHIFTARRAY ]
	TEST GETARRAY 0 moreSpirits ->STRING [ rum ] ->STRING [ UNSHIFTARRAY ]


	: moreSpirits => UNSHIFTARRAY ->STRING [ rum ] spirits

	TEST GETARRAY 0 spirits ->STRING [ vodka ] ->STRING [ ARRAYPUSH and GETARRAY ]

	REBIND spirits => ARRAYPUSH spirits ->STRING [ vodka ]

	: spirits => NEWARRAY

]


TESTBLOCK [

	TEST GETHASH pages bookdata 	132 	->STRING [ Hash clone functionality ]

	: newbookdata => HASHSET bookdata pages 10

	TEST GETHASH pages bookdata 132 ->STRING [ HASHSET ]

	: bookdata => HASHSET NEWHASH pages 132

	]

	TESTBLOCK [

	TEST GETHASH pages bookdata 134 ->STRING [ get and set hashes ]



	REBIND bookdata => SETHASH pages 134 bookdata

	BIND bookdata => NEWHASH

]

TESTBLOCK [


TEST  tests: GETHASH AKEY TOK TESTHASH expects: 1  [ HASH STRING EQUALITY ]


REBIND TESTHASH => SETHASH  AKEY TOK 1 TESTHASH


: TESTHASH => NEWHASH

]


DEFINE HASHSET => [ SETHASH ROLL 2 ROLL 2 ]


REBIND GETHASH => ->FUNC [
                   IF EQUAL HASH TOK  GETTYPE A-HASH
                      [  REAL-GETHASH A-KEY A-HASH ]
                      [ EXIT PRINTLN A[ at line LOCATIONOF A-HASH ]A EMIT A-HASH EMIT SPACE EMIT [ GETHASH ERROR: EXPECTED HASH BUT GOT ]  ]
                      ARG A-HASH =>
                      ARG A-KEY  => ->STRING
                   ]
ALIAS REAL-GETHASH => GETHASH TOK


DEFINE HASHGET => [ GETHASH SWAP ]



DEFINE returns => [ ]
DEFINE expects: => [ ]
DEFINE tests: => [ ]


DEFINE )- => [ )- TOK ]




DEFINE -( => [
	IF  NOT EQUAL )- TOK A
	[ EXIT DUMP A  EMIT [ UNBALANCED AT ] ]
	[  ]

	ARG A =
]


DEFINE LEXLOOKUP => [ GETFUNCTION ]

DEFINE = => [ EXIT PRINTLN [ = HAS BEEN DEPRECATED ] ]
ALIAS LAMBDA => ->LAMBDA TOK
ALIAS CODE => ->CODE TOK
ALIAS ->STRING => ->STRING TOK
ALIAS IT =>  DUP TOK
ALIAS NEVAL =>  NATIVE_EVAL TOK
ALIAS => => TOK TOK
ALIAS => TOK TOK TOK
DEFINE ( TOK [ ]
DEFINE ) TOK [ ]

TESTBLOCK [
	TEST STRING-CONCATENATE A TOK B TOK AB TOK ->STRING [ STRING-CONCATENATE ]
]
DEFINE _ TOK  [ STRING-CONCATENATE STRING-CONCATENATE SWAP SPACE ]

DEFINE CD TOK [ _ SWAP ]

DEFINE LS TOK MACRO [ DIRECTORY-LIST CWD ]

TESTBLOCK [
	TEST  ./  CWD  [ CWD ]
]


DEFINE CWD TOK MACRO [ CURRENT-DIRECTORY ]

DEFINE PWD TOK MACRO [ PRINTLN CURRENT-DIRECTORY ]

: CURRENT-DIRECTORY TOK ->STRING [ ./ ]



DEFINE PROMPT_HOOK TOK [ PRINTLN DUP EMIT > NEWLINE PRINTLN  [ throff interactive mode  ]  ]
NAMETESTBLOCK [ Testing PICK ] [
	TEST DROP SWAP DROP SWAP PICK 1 A TOK B TOK B TOK [ PICK ]
]

NAMETESTBLOCK [ Testing basics ] [
	TEST  ->STRING COMPVAL TOK COMPVAL TOK  [ String/token equality ]
	TEST 1 0  [  test failure  ]
	TEST 1 1  [  test itself  ]
]

DEFINE STATEMENT TOK [

	IF EQUAL MARKER MARKERT [ ]
	[
		EXIT .S PRINTLN ->STRING [ altered stack ] EMIT SPACE EMIT GETFUNCTION NAME TOK PRINTLN ->STRING [ ===ERROR===   ] EMIT LOCATIONOF NAME
	]

	: MARKERT =>
	CALL TESTS
	MARKER


	: MARKER TOK STRING-CONCATENATE STACKMARKER: TOK ->STRING GETFUNCTION TESTS TOK

	: TESTS TOK THIN
	: NAME TOK
]

DEFINE NAMETESTBLOCK [ IF RUNTESTS [ TESTBLOCK PRINTLN ] [ DROP DROP ] ]

DEFINE TESTBLOCK TOK [
	IF RUNTESTS [
		IF EQUAL MARKER MARKERT [ ]
		[
			EXIT .S PRINTLN LOCATIONOF TESTS EMIT [ ERROR: test altered stack at: ]

		]

		BIND MARKERT TOK
		CALL   TESTS

		MARKER


		: MARKER TOK STRING-CONCATENATE STACKMARKER:  TESTS
	]
	[ ]

	: TESTS TOK
]

DEFINE TEST TOK [

		IF EQUAL EXPECTED TESTVAR
		[ PRINTLN ->STRING [ ... Test passed ] ]
		[
			PRINTLN TESTVAR  EMIT SPACE EMIT [ but got ]
			EMIT EXPECTED EMIT SPACE  EMIT [ Expected ]
			EMIT [ ... Test failed ... ]
		]

		EMIT DESCRIPTION-TEXT



		EMIT SPACE
		EMIT ->STRING [ Testing ]

		EMIT : TOK
		EMIT LOCATIONOF DESCRIPTION-TEXT

	ARG DESCRIPTION-TEXT TOK ->STRING
	ARG EXPECTED TOK
	ARG TESTVAR TOK
]

: DESCRIPTION TOK ->STRING [ SHOULD BE OVERWRITTEN ]




DEFINE UNFUNC TOK [
COMMENT [
	IF EQUAL FTYPE CODE TOK
	[ ->LAMBDA GETFUNCTION F TOK ]
	[ MACRO [ F ] ]


	DROP [ PRINTLN FTYPE EMIT [ TYPE IS: ] ]
	BIND FTYPE => GETTYPE GETFUNCTION F TOK

	ARG F TOK
	]

]


ALIAS STRING! TOK ->STRING TOK
DEFINE ->ARRAY TOK [ SETTYPE ARRAY TOK ]
DEFINE ->CODE TOK [ SETTYPE CODE TOK ]
DEFINE ->STRING TOK [ SETTYPE STRING TOK ]
DEFINE ->LAMBDA TOK [ SETTYPE LAMBDA TOK ]



ALIAS ARG TOK : TOK
ALIAS BIND TOK : TOK
: DEFINE TOK MACRO  ->FUNC [

	SCRUBLEX DEFINE_F TOK
	SCRUBLEX DEFINE_N TOK

	IF NOT OR
			EQUAL LAMBDA TOK GETTYPE GETFUNCTION  DEFINE_F TOK
			EQUAL CODE TOK GETTYPE GETFUNCTION  DEFINE_F TOK
				MACRO [ EXIT
				PRINTLN GETTYPE  DEFINE_F TOK  EMIT TYPE: PRINTLN SPACE
				PRINTLN GETFUNCTION  DEFINE_F TOK  EMIT VALUE: PRINTLN DEFINE_N EMIT [ ATTEMPTED TO BIND SOMETHING THAT IS NOT A FUNCTION TO: ]  ]
				MACRO [ BIND DEFINE_N ->FUNC GETFUNCTION DEFINE_F TOK  ]
		COMMENT [ PRINTLN DEFINE_N EMIT [ DEFINING:  ] ]
		: DEFINE_F TOK
		: DEFINE_N TOK


	]

: OR TOK  ->FUNC [
	IF A
	[ TRUE ]
	[ IF B
		[ TRUE ]
		[ FALSE ]
	]

	: A TOK : B TOK
]


ALIAS ADDWORD TOK : TOK
: ALIAS TOK MACRO  ->FUNC [ SETLEX SWAP GETFUNCTION SWAP ]

: CHAT TOK  ->FUNC [ PRINTLN ]
COMMENT [ BASIC FUNCTIONS DEFINED! ]

: PRINT3 TOK ->FUNC  [ EMIT EMIT EMIT ]
: PRINT2 TOK ->FUNC  [ EMIT EMIT ]
: DO TOK ->FUNC  [ NATIVE_EVAL ]

: CRLF TOK ->FUNC [ STRING-CONCATENATE CR NEWLINE ]
: CR TOK ->FUNC  [ NUM2CHAR 10   ]
: NEWLINE TOK ->FUNC  [ NUM2CHAR 13  ]
: 2DUP TOK ->FUNC  [ OVER OVER ]
: 2DROP TOK ->FUNC  [ DROP DROP ]
: TUCK TOK ->FUNC  [  OVER SWAP ]
: NIP TOK ->FUNC  [ DROP SWAP ]
: ROT TOK ->FUNC  [ ROLL 2  ]
: SWAP TOK ->FUNC  [ ROLL 1 ]
: OVER TOK ->FUNC  [ PICK 1 ]

: DUP TOK ->FUNC [  PICK ZERO ]
COMMENT [ LOADING STANDARD LIBRARY .... ]
: COMMENT TOK ->FUNC [ DROP ]
: ->FUNC TOK SETTYPE CODE TOK [ SETTYPE CODE TOK ]
: FALSE TOK EQUAL 0 1
: TRUE TOK EQUAL 1 1




: RUNTESTS TRUE


ITROFF
IDEBUGOFF
DEBUGOFF
 ->FUNC [ EXIT .S PRINTLN [ ERROR: Read past stack bottom attempted ]  ]


`

return str
}
