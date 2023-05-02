# Bits and Bytes

The *byte* is generally the smallest item of storage that is directly
accessible in modern computers, and symbiotically the byte is the
smallest item of storage used by modern programming languages. We could think of the byte as the atomic level of data.

But there is a subatomic story. A byte consists of eight bits, where a
*bit* is the electronic equivalent of a bivalent entity. In other words,
a bit is capable of being in exactly two states. We can think of these
two states as true/false, +1/-1, or 1/0. The machine has electronic
representations of bits, in groups of eight, and makes these groups of
eight bits directly accessible as bytes. A byte, being the atomic data
type, is the space used to store `char` data. You learn a lot more about this in our course CDA 3100 *Computer Organization I*.

One of the many legacies of C, inherited by C++, is an ability to access
subatomic data, data at the bit level, indirectly through *bit
operators*. Moreover, all modern general-purpose computers have hardware support for these bit operators, which makes them extremely fast. Many C programmers use bit operators to implement various bit-level access utilities, the most popular being the *bitvector*. The C++ class allows us to descend to the bit level carefully and then encapsulate this sub-atomic programming behind a public interface. We will illustrate this entire process later in the chapter by constructing a bitvector class, a class that stores and retrieves vectors of bivalent values at the bit level.

The byte can store any eight-digit unsigned binary number (the binary
digits are usually called bits). The range of such numbers begins with
0~(10) = 00000000~(2)~ and ends with 11111111~(2)~ = 255~(10)~. Here is
the range of numbers representable in one byte, expressed in various
notations:

*Name*

``*Base*

``*1-Byte Range*

Binary

2

```00000000 … 11111111`

Octal

8

```000 … 377`

Decimal

10

```0 … 255`

Hexadecimal

16

```00 … FF`

The hexadecimal, or hex, representation is perhaps the most convenient
to use for byte values, since it consists of exactly two hex digits. The
hex digits are ` 0 1 2 3 4 5 6 7 8 9 A B C D E F ` and the range of
two-digit hex numbers corresponds to the decimal range 0 ... 255 =
16^2^ - 1.

For further reading, see *Number Systems*, Appendix C of Dietel.\

[]{#Link2}

# Bit Operators in C/C++

The following table summarizes the bit operators C++ inherits from C:

  --------------------- ------------ --------------- --------------------- ------------------------------
  *operation*``     *symbol*   *type*         *infix version*      *accumulator version* ``
   and                  `&`        binary         `z = x & y`          `z &= y`
   or                   `|`        binary         `z = x | y`          `z |= y`
   xor                  `^`        binary         `z = x ^ y`          `z ^= y`
   not                  `~`        unary          `z = ~y`             (na)
   left shift           `<<`       binary         `z = y << n`         (na)
   right shift ``    `>>`       binary ``   `z = y >> n` ``   (na)
  --------------------- ------------ --------------- --------------------- ------------------------------

The first four bit operators listed in the table are based on the AND,
OR, XOR, and NOT operators for single bits, given by the following
tables:

Such tables are often called *truth tables*, a term from symbolic logic
for the table of truth values of a logical expression based on
truth/falsity of the variables in the expression. This is one of the
areas where logic, math, electrical engineering, and computer
architecture intersect. The table represents: (1) all possible values of
a logical expression (*logic*), (2) a table of values for a binary
function (*math*), (3) a table of electrical charge input/output for a
circuit (*electrical engineering*), and (4) a basic component of a
digital computer (*computer architecture*). In the modern computer,
these bit-level operations have one-clock-cycle hardware
implementations.

The C/C++ bit operators `&`, `|`, `^`, and `~` are defined for any
integral type (such as `char`, `short`, `size_t`, and `long`) by
applying the single-bit operations *bitwise*. In other words, these
operators simply apply the single-bit operations in parallel for each
bit of the operand. The parallel application of the single-bit
operations is also supported by hardware, so that the bitwise operators
also operate in one clock cycle \-- fast.

The other two bit operators shown are the left shift (`operator << ()`)
and right shift (`operator >> ()`). These two-argument operators take an
integral type as the left operand and a non-negative integer as the
right operand, and return the same integral type as the left operand.
The prototypes for the shift of unsigned long, for example, are:

      unsigned long operator << (unsigned long , int);
      unsigned long operator >> (unsigned long , int);

The shift operators move the bits of the left argument the number of
times given by the right argument and return the (shifted) left argument
as a value. Bits that are shifted out of the left operand are lost (it
is said that these bits go into the \"bit bucket\") and bits uncovered
by shifting are replaced by bit `0`. Left and right shift also have
hardware support, so that they execute much more quickly than standard
arithmetic operators.

## Notes:

1.  A point not to be confused over is the term \"binary\" as it is
    often applied to operators. A *binary operator* is an operator that
    takes two arguments. Similarly, a *unary operator* is an operator
    that takes only one argument. In this usage, the term \"binary\" is
    referring to the number of arguments, not their type.
2.  Left shift `<<()` and right shift `>>()` are native C/C++ operators.
    The C++ input and output operators are overloads of these natives of
    C.
3.  Shift operators are generally recommended only for unsigned integer
    types. Some shift operator behaviors on signed type are machine
    dependent, making their use unpredictable.
4.  The shift operators are undefined when the right argument is
    negative or is larger than the number of bits in the binary
    representation of the first argument.

[]{#Link3}

# Examples of Bit Operator Calculations

This table shows some examples of bit operations in action.

You should be able to perform such calculations yourself. Some practice
might be called for to ensure this.\

[]{#Link4}

# Defining Class BitVector

Our goal in the remainder of this chapter is the development of a bit
vector class, which we will call `BitVector`. Following the established
pattern, we need a *public interface* and an *implementation plan*.

## BitVector Public Interface

There are not that many things one can do to a bit: a bit can be *set*
(meaning give it the value `1`) or *unset* (meaning give it the value
`0`); and one can *flip* a bit, (meaning change its value from `0` to
`1` or from `1` to `0`). An ability to *test* the value of a bit is
necessary if access is to be useful. The return value of a test needs to
be an integer type, which we will interpret as a bit value. These four
functions form the core of the user interface for our BitVector class.

    namespace fsu
    {
    class BitVector
    {
      public:
        ...
        void Set   (size_t index);        // make index bit = 1
        void Unset (size_t index);        // make index bit = 0
        void Flip  (size_t index);        // change index bit 
        int  Test  (size_t index) const;  // return index bit value
        ...
    }; // class BitVector
    } // namespace fsu

Note that we have placed the `BitVector` class in the `namespace fsu`.

The public interface should also include constructors, a destructor,
assignment operator, and a method `Size()` which is intended to return
the number of bits stored in a BitVector object. (Due to implementation
details, the unsigned integer `Size()` will always be a multiple of
eight.) A design decision (that could be reversed later, without harm to
existing client programs) is that class `BitVector` will not have a
default constructor: The size requirement (number of bits) of a
BitVector object must be known in advance and passed to the constructor
as a parameter.

    class BitVector
    {
      public:
        ...
        explicit BitVector  (size_t);     // construct a BitVector with specified size
                 BitVector  (const BitVector&); // copy constructor      
                 ~BitVector ();           // destructor
        BitVector& operator = (const BitVector& a);  // assignment operator
        size_t Size   () const;           // return size of bitvector
        ...
    };

We will also overload the `Set()`, `Unset()`, and `Flip()` methods so
that they apply to the entire `BitVector` object when invoked without an
index parameter:

    class BitVector
    {
      public:
        ...
        void Set   ();         // make all bits = 1
        void Unset ();         // make all bits = 0
        void Flip  ();         // change all bits 
        ...
    };

These components are collected into a summary at the end of this
chapter.\

[]{#Link5}

# BitVector Implementation Plan

The implementation plan for BitVector is clever, embodying many of the
bit manipulation tricks invented and used by C programmers over the
years. Don\'t feel inadequate if you didn\'t immediately conceive a good
implementation plan. But be certain you understand the plan, as well as
the details of the implementation itself. You will need to demonstrate
and use this knowledge several times in this course as well as in
succeeding courses and your professional life.

The basic storage facility for bits will be a vector of bytes. A vector
of *n* bytes contains 8*n* bits (hence the size of a BitVector object is
a multiple of 8). The challenge is to index and access at the bit level.
The math of this is fairly straightforward: to get to the index *k* bit,
we divide *k* by 8 to get the byte and then count up to the bit in that
byte corresponding to the remainder. Here is an example:

Suppose we have a array v of bytes and we want the 29th bit. There are 8
bits in v\[0\], 8 bits in v\[1\], and so on. Indexing these bits
(beginning with index 0), bits 0 through 7 are in byte v\[0\], bits 8
through 15 are in byte v\[1\], bits 16 through 23 are in v\[2\], and
bits 24 through 31 are in v\[3\]. Therefore, bit 29 is bit 5 in v\[3\].
Note that 29/8 = 3 (the byte number) and 29%8 = 5 (the bit number).

To get to the bit with index *k*, we first divide *k* by 8 (integer
division) obtaining a quotient *q* and remainder *r*. (That is, *k* =
8\**q* + *r*.) Then bit\[*k*\] is the *r*th bit of the *q*th byte, i.e.,
bit *r* of v\[*q*\]. We have now reduced the challenge to two problems:

1.  Find the byteArray\_ index for bit index *k* (the quotient)
2.  Access the bit for bit-index *k* (the remainder)

Problem 1 is essentially solved, just return the value *k*/8. To solve
problem 2 we must finally \"byte the bullet\" and do some subatomic
programming using the concept of *mask*.

## The Mask

A mask for a pattern of bit locations in an integral type is a word of
that type whose binary representation has bit `1` in the specified
locations and bit `0` elsewhere. For our purposes here, we may as well
assume the type is `char`, so that a word consists of a single byte, or
eight bits. (Alternative approaches replace bytes with 32-bit words, and
use masks for `unsigned long`, a 32-bit word.) Assuming type `char`,
there are eight possible 1-bit masks, as follows:

    00000001 = 00000001 << 0    // == 1   == 0x01
    00000010 = 00000001 << 1    // == 2   == 0x02
    00000100 = 00000001 << 2    // == 4   == 0x04
    00001000 = 00000001 << 3    // == 8   == 0x08
    00010000 = 00000001 << 4    // == 16  == 0x10
    00100000 = 00000001 << 5    // == 32  == 0x20
    01000000 = 00000001 << 6    // == 64  == 0x40
    10000000 = 00000001 << 7    // == 128 == 0x80

Notice that each of these can be obtained from 00000001~(2)~ = 1~(10)~ =
0x01~(16)~ by applying the left shift operator, as indicated in the list
above. (The decimal and hexadecimal representations of the mask are also
shown above.)

A mask can be used to access an individual bit of a byte. Suppose, for
example, that we want to access the third bit from the right of byte
`x`. Calculate the value

      y = 4 & x
        = 00000100 & x

Note that `y` can have only two values:

      00000100 = 4, or
      00000000 = 0.

The test `(y != 0)` and the third bit of `x` have essentially the same
value, true or false. In effect, the value of the third bit (bit index
2, relative to the byte) of `x` is equal to the value returned by the
boolean expression `(((1 << 2) & x) != 0)`. The mask `(1 << 2)` allows
us to access the bit value indirectly using a test. The access is also
very fast, because all of the operations have direct hardware support.

We have now solved problem 2. To access bit index `k`, we use
`Mask = 1 << (index % 8)` applied to the byte
`byteArray_[ByteNumber (index)]`. This completes the implementation plan
for BitVector.

The following private section of the BitVector class facilitates this
implementation plan:

    class BitVector
    {
        ...
      private:
        // data
        unsigned char * byteArray_;
        size_t          byteArraySize_;
>
        // methods
        size_t        ByteNumber (size_t indx) const;
        unsigned char Mask       (size_t indx) const;
    };

The two private methods `ByteNumber()` and `Mask()` capture the two main
ideas for locating a particular bit in the bit vector.\

[]{#Link6}

# Implementing BitVector

Implementations of the `BitVector` member functions are not lengthy, and
for the most part are straightforward applications of the ideas
expressed above combined with appropriate use of the bit operators
discussed earlier. Exceptions are the two private methods, whose
implementations use clever ideas for optimizing the computations that
make them a \"bit\" difficult to comprehend. (This is one good reason
for building the BitVector class, so that these precious details can be
committed to computer, rather than human, memory.)

## Class Definition

Here for reference is the complete class definition:

    class BitVector
    {
      public:
        explicit BitVector  (size_t);     // construct a BitVector with specified size
                 BitVector  (const BitVector&); // copy constructor      
                 ~BitVector ();           // destructor
>
        BitVector& operator = (const BitVector& a);  // assignment operator
>
        size_t Size () const;             // return size of bitvector
>
        void Set   (size_t index);        // make index bit = 1
        void Set   ();                    // make all bits = 1
        void Unset (size_t index);        // make index bit = 0
        void Unset ();                    // make all bits = 0
        void Flip  (size_t index);        // flip index bit (change value of bit)
        void Flip  ();                    // flip all bits 
        int  Test  (size_t index) const;  // return index bit value
       
      private:
        // data
        unsigned char * byteArray_;
        size_t          byteArraySize_;
>
        // methods
        size_t               ByteNumber (size_t indx) const;
        static unsigned char Mask       (size_t indx) const;
    };

## BitVector::ByteNumber

The implementation of `ByteNumber()` is a good example of \"clever\". We
have already established that the private method `ByteNumber(index)`
should return `index/8`, but the code below seems to return something
different: `index >> 3`. This is actually an optimization that executes
faster while accomplishing the same result. The critical observation is
that 8 = 2^3^, and for powers of 2 the right shift operator achieves the
same result as integer division: each right shift has the same effect as
division by 2, so right shifting 3 is the same as division by 8. (The
reader should verify this using pencil and paper.) But, because right
shift has hardware support, it is much faster than integer division,
hence, the optimization.

    size_t BitVector::ByteNumber (size_t index) const
    {
      // return index / 8
      // shift right 3 is equivalent to, and faster than, dividing by 8
      index = index >> 3;
      if (index >= byteArraySize_)
      {
        std::cerr << "** BitVector error: index out of range\n";
        exit (EXIT_FAILURE);
      }
      return index;
    }

## BitVector::Mask

A similarly clever trick is used in to implement the `Mask` method. This
time, the observation is that the remainder when dividing a number by a
power k of 2 is equal to the last k bits of the number in binary
representation. (These are the bits that are lost to the bit bucket
during the right shift by k.) Thus, the remainder when dividing by 8 is
just the last 3 bits of the number, which we can access by bitwise AND
with a mask of 7 = 0x07 = 00000111~2~. Again, using the mask and the AND
operator is much faster than division because of the hardware support
for the bit operators. Then, we shift 00000001~2~ = 0x01 = 1 by that
amount to have the mask for `index`.

    unsigned char BitVector::Mask (size_t index) const
    {
      // return mask for index % 8
      // the low order 3 bits is the remainder when dividing by 8
      size_t shiftamount = index & 0x07;  // low order 3 bits
      return 0x01 << shiftamount;
    }

**Note:**The `Mask` method is declared as `static`. The connotation is
that the method does not use any (non-static) class variables. It is
good software engineering practice to declare any method that does not
need access to object-level variables `static`. This results in improved
compiler optimizations and provides protection against inadvertant
access to object data in the implementation of the method.

Here is a picture of the bitvector, layed out with the byteVector
elements right to left so that the order of the bits increases from
right to left (like a number):

    i = ... 76543210 98765432 10987654 32109876 54321098 76543210
    b = ... 5        4        3        2        1        0        
    v = ... ******** ******** ******** ******** ******** ********
>
     where i = bit index   (showing only the first digit of i)
           b = byte number (aka byteVector index)
           v = byteArray   (spaces just for readability)

So for example if `i` = 19 then

     i = 19
     b = 19/8 = 2 and
     m = (0x01 << 19%8) = (0x01 << 3) = 00001000,

and the specific picture is:

    b = ... 5        4        3        2        1        0        
    v = ... ******** ******** ******** ****x*** ******** ********
    i = ...                                1            
    m =                                00001000                  

In words: bit 19 is in the array element 2 and has mask `00001000`. Note
that `mask & v[2] = 0000x000`: the byte `0000x000` is zero iff the bit
`x` is zero, which is the way Test(i) is implemented:

     Test(11) = 1 if x = one and
     Test(11) = 0 if x = zero

Thus we have a way of detecting the value of bit `11` without having
direct access to it.

## BitVector Constructor

The constructor takes a size parameter and calls operator
`new unsigned char` with parameter equal to the integer that is just
larger or equal to `numbits/8`.

    BitVector::BitVector (size_t numbits)  // constructor
    {
      byteArraySize_ = (numbits + 7)/8;
      byteArray_ = new unsigned char [byteArraySize_];
      if (byteArray_ == 0)
      {
        std::cerr << "** BitVector memory allocation failure -- terminating program.\n";
        exit (EXIT_FAILURE);
      }
      for (size_t i = 0; i < byteArraySize_; ++i)
        byteArray_[i] = 0x00;
    }

## BitVector Assignment Operator

The assignment operator follows the typical pattern:

    BitVector& BitVector::operator = (const BitVector& bv)  //
    assignment operator
    {
      if (this != &bv)
      {
        if (byteArraySize_ != bv.byteArraySize_)
        {
          delete [] byteArray_;
          byteArraySize_ = bv.byteArraySize_;
          byteArray_ = new unsigned char [byteArraySize_];
          if (byteArray_ == 0)
          {
            std::cerr << "** BitVector memory allocation failure -- terminating program.\n";
            exit (EXIT_FAILURE);
          }
        }
        for (size_t i = 0; i < byteArraySize_; ++i)
          byteArray_[i] = bv.byteArray_[i];
      }
      return *this;
    }

## BitVector API

The BitVector API consists of the public methods `Set`, `Unset`, `Flip`,
and `Test` (two versions of each of the first three).

The `Test(index)` method, like the remaining methods taking an index
parameter, is a one-liner. But the line is\...clever. The idea is to
isolate the indexed bit with a mask and test for non-zero result, as we
established earlier. Look carefully at the computation shown and be
certain you understand it.

    int BitVector::Test  (size_t index) const  
    // return specified bit value
    {
       return 0 != (byteArray_[ByteNumber(index)] & Mask(index));
    }

Our final example is the `Set(index)` method, another one-liner:

    void BitVector::Set (size_t index)
    // set specified bit
    { 
      byteArray_[ByteNumber(index)] |= Mask(index);                                                
    }

The rest of the implementation of BitVector is left as an assignment.
Enjoy.

\

[]{#Link7}

# Sample BitVector Client

Appended below is source code for a straightforward functionality test
client for BitVector. Here are some key attributes of the program:

-   The program allows the user to set the bitvector size at the
    beginning of the test.
-   To allow the user to set the bitvector size, dynamic allocation must
    be used (operators `new` and `delete`).
-   The user is presented a simple interface that gives access to most
    of the public interface of BitVector.
-   Assignment and copy constructor are not tested by this program.

The test client can also be used as a device to discover by experiment
how bitvectors behave. An executable of this client may be accessed in
`area51`. If you want to get the feel of bitvector use, give this a try.
Copy the executable into your directory and run it by typing the name.

    /*  fbitvect.cpp
>
        testing BitVector
    */
>
    #include <xbitvect.cpp>
>
    void display_menu(unsigned int size);
>
    int main()
    {
      char selection;
      size_t size, index;
      std::cout << "Welcome to fbitvect. Enter size of BitVector: ";
      std::cin >> size;
      fsu::BitVector * bvptr = new fsu::BitVector (size);
      display_menu(size);
      do
      {
        std::cout << "  Enter [op][index] (m for menu, x to exit): ";
        std::cin >> selection;
        switch(selection)
        {
          case 'S':
            bvptr->Set();
            break;
          case 'U':
            bvptr->Unset();
            break;
          case 'F':
            bvptr->Flip();
            break;
          case 's':
            std::cin >> index;
            bvptr->Set(index);
            break;
          case 'u':
            std::cin >> index;
            bvptr->Unset(index);
            break;
          case 'f':
            std::cin >> index;
            bvptr->Flip(index);
            break;
          case 't': case 'T':
            std::cin >> index;
            std::cout << "  v[" << index << "] == ";
            if (bvptr->Test(index))
              std::cout << "1\n";
            else
              std::cout << "0\n";
            break;
          case 'd': case 'D':
            std::cout << "  v == " << *bvptr << '\n';
            break;
          case 'm': case 'M':
            display_menu(size);
            break;
          case 'x': case 'X':
            break;
          default:
            std::cout << "  command not found\n";
        }
      }
      while (selection != 'x' && selection != 'X');
      delete bvptr;
      std::cout << "Thank you for testing BitVector\n";
      return 0;
    }
>
    void display_menu(unsigned int size)
    {
       std::cout << "     BitVector (" << size << ") v;\n"   
                 << "     operation                                entry\n"
                 << "     ---------                                -----\n"
                 << "     v.Set            ()  ......................  S\n"
                 << "     v.Set            (index)  .................  s\n"
                 << "     v.Unset          ()  ......................  U\n"
                 << "     v.Unset          (index)  .................  u\n"
                 << "     v.Flip           ()  ......................  F\n"
                 << "     v.Flip           (index)  .................  f\n"
                 << "     v.Test           (index)  .................  t\n"
                 << "     cout << v  ................................  d\n"
                 << "     display this Menu  ........................  m\n"
                 << "     eXit program  .............................  x\n";
    }

This harness does not test the copy constructor and assignment operator.
To expand the program to include test of the assignment operator would
require defining three bitvector objects (instead of one as above) and
adding for each of these objects all of the options above plus
assignment plus three-way assignment. A function call (passing a
bitvector by value) and return would be a good test of the copy
constructor.
