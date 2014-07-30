{-# LANGUAGE LambdaCase #-}
import Prelude hiding (Left, Right)
import Control.Applicative
import Control.Monad.State
import Data.List
import qualified Data.Map as M
import Data.Maybe
import Data.Ord hiding (Down)

type Position = (Int,Int)
type Level = M.Map Position Cell

data Cell = Cell Tile Object deriving (Show, Eq, Ord)
data Tile = RedTile | SolidTile | CrackedTile | NoTile deriving (Show, Eq, Ord)
data Object = Taco | BlueBlock | RedBlock | TrelloCoin | NoObject deriving (Show, Eq, Ord)

type Path = [Direction]
data Direction = Up | Down | Right | Left deriving Show

directions :: [Direction]
directions = [Up, Down, Right, Left]

findTaco :: Level -> Position
findTaco = head . M.keys . M.filter hasTaco
  where
    hasTaco (Cell _ Taco) = True
    hasTaco _             = False

countCoins :: Level -> Int
countCoins = M.fold addCoin 0
  where 
    addCoin (Cell _ TrelloCoin) = (+1)
    addCoin _                   = id

solveLevel :: Level -> Path
solveLevel = reverse . head . sortBy (comparing length) . solveLevel' [] [] []

solveLevel' :: [Level] -> Path -> [Path] -> Level -> [Path]
solveLevel' seen path acc level
  | elem level seen       = []
  | countCoins level == 0 = path : acc
  | otherwise             = concatMap (solveLevel'' (insert level seen) path acc level) directions 

solveLevel'' :: [Level] -> Path -> [Path] -> Level -> Direction -> [Path]
solveLevel'' seen path acc level dir =
  concatMap (solveLevel' seen (dir:path) acc) $ maybeToList (move level dir)

move :: Level -> Direction -> Maybe Level
move level dir = moveFrom level >>= pushBlock >>= moveTo
  where 
    pos = findTaco level
    nextPos = nextPosition pos dir
    secondNextPos = nextPosition nextPos dir
    -- Move off the current tile.
    moveFrom level = do
      cell <- fmap stepOff $ M.lookup pos level
      return $ M.insert pos cell level
      where
        stepOff (Cell tile _) = case tile of
          SolidTile   -> Cell CrackedTile NoObject
          CrackedTile -> Cell NoTile NoObject
          _           -> Cell tile NoObject
    -- Try to push a block out of the way, if there is one.
    pushBlock level = M.lookup nextPos level >>= \case
      (Cell _ RedBlock) -> case M.lookup secondNextPos level of
        Just (Cell NoTile _)   -> Just $ M.insert secondNextPos (Cell RedTile NoObject) level
        Just (Cell t NoObject) -> Just $ M.insert secondNextPos (Cell t RedBlock) level
        _                      -> Nothing
      (Cell _ BlueBlock) -> Nothing
      _                  -> Just level
    -- Move onto the next tile.
    moveTo level = M.lookup nextPos level >>= \case
      (Cell NoTile _) -> Nothing
      (Cell t _)      -> Just $ M.insert nextPos (Cell t Taco) level


nextPosition :: Position -> Direction -> Position
nextPosition (r,c) = \case
  Up    -> (r-1,c)
  Down  -> (r+1,c)
  Right -> (r,c+1)
  Left  -> (r,c-1)

parseLevel :: [String] -> Level
parseLevel = parseLevel' 0

parseLevel' :: Int -> [String] -> Level
parseLevel' _ []     = M.empty
parseLevel' r (x:xs) = M.union (parseRow r 0 x) (parseLevel' (r+1) xs)

parseRow :: Int -> Int -> String -> Level
parseRow _ _ []       = M.empty
parseRow r c (x:y:zs) = M.union cell $ parseRow r (c+1) zs
  where cell = M.singleton (r,c) $ parseCell x y

parseCell :: Char -> Char -> Cell
parseCell x y = Cell (parseTile x) (parseObject y)
  where
    parseTile = \case
      'X' -> RedTile
      '=' -> SolidTile 
      '-' -> CrackedTile
      ' ' -> NoTile
      _   -> error "unable to parse tile"
    parseObject = \case
      'P' -> Taco
      '#' -> BlueBlock
      '^' -> RedBlock
      'J' -> TrelloCoin
      ' ' -> NoObject
      _   -> error "unable to parse object"

levels :: [Level]
levels = map parseLevel
  -- Level 1
  [["=#=#=#"
   ,"XPX XJ"
   ,"= =^= "]
  -- Level 2
  ,["-P- = = "
   ,"      = "
   ,"-J- -J- "]
  -- Level 3
  ,["-P-^  =J"
   ,"-^    -^"
   ,"-#      "
   ,"XJX X -J"]
  -- Level 4
  ,["=JXJ-   X^"
   ,"X   -J  XJ"
   ,"= X^  XJ-J"
   ,"-^=#=^-^-P"
   ,"    X   -J"]
  -- Level 5
  ,["= - -P= =#"
   ,"-J=^XJX^=#"
   ,"XJX^X^- =#"
   ,"=#=#XJ  =J"
   ,"= -^=J  =J"]
  -- Level 6
  ,["= =   XJ= "
   ,"-J- XJ-^-^"
   ,"XJX =^XJ=J"
   ,"-^-^X^X = "
   ,"=^=PX^=#X^"]
  -- Level 7
  ,["XJ=   =J- "
   ,"= -   -J-J"
   ,"=^-J  X =^"
   ,"=P-^XJXJ=^"
   ,"= - -^= - "]
  -- Level 8
  ,["=^X -J-   "
   ,"-J- =#-JXJ"
   ,"-^=J-J=#- "
   ,"  =^-J-^X "
   ,"  = = -PX^"]
  -- Level 9
  ,["X = -^XJX^"
   ,"-^X^X =J  "
   ,"- =#X   XJ"
   ,"- =P= = - "
   ,"=^X^-^-^-J"]
  -- Level 10
  ,["=#  =J-JXP"
   ,"  X -^- -^"
   ,"XJ=#-   - "
   ,"- X^=J  = "
   ,"=JX^  - = "]
  -- Level 11
  ,["= =   -J- "
   ,"X^=#=#  X "
   ,"-^-^-J-J= "
   ,"=#-^=^X - "
   ,"=J-PXJXJ-J"]
  -- Level 12
  ,["=^  =   X^"
   ,"= = =^-^XJ"
   ,"XJ- X -   "
   ,"-PX -^=#-J"
   ,"=^-J=JX^= "]
  -- Level 13
  ,["X - = = X "
   ,"X =J-^XJ= "
   ,"-^=^X^-J=^"
   ,"- XJ=#-^=J"
   ,"-^-PXJ-J=#"]
  -- Level 14
  ,["=JX^=#-^X^"
   ,"XJ=#-^= - "
   ,"- -JXJ= X^"
   ,"  -PX^X   "
   ,"=JX =^-J-^"]
  -- Level 15
  ,["=JXJ-^X   "
   ,"-^=J=^-JX^"
   ,"-P-J=#X^=#"
   ,"-^= =^X - "
   ,"= XJ=#-J=^"]
  -- Level 16
  ,["-^= =J  - "
   ,"-^X - XJ-J"
   ,"=J=^X^=#= "
   ,"=^=P-^=JX^"
   ,"-J  -^XJ= "]
  -- Level 17
  ,["X =J= =^= "
   ,"-J-^=^X^X^"
   ,"-P=J=#- -J"
   ,"X X^- - -^"
   ,"  X -J  X "]
  -- Level 18
  ,["- X = XJ-^"
   ,"X -^=#-^=#"
   ,"X XJ-J-   "
   ,"=J  =#XJ-J"
   ,"  -^=P-^=J"]
  -- Level 19
  ,["XJX^-J=#=^"
   ,"- = =^=J=^"
   ,"-^  =#X XP"
   ,"-^=^=J  = "
   ,"- -^=J-JXJ"]
  -- Level 20
  ,["-J-J= =J  "
   ,"- -J=#- X "
   ,"-^= X -^=P"
   ,"= =^=#-J=J"
   ,"  -J- -^  "]
  -- Level 21
  ,["- -J=^  =#"
   ,"  XJ-^XJ-^"
   ,"X^  -^X^  "
   ,"- = = X -^"
   ,"XJ-^X^=#XP"]
  -- Level 22
  ,["-J  =JX =P"
   ,"=#XJ-JX =^"
   ,"  -^X =J=#"
   ,"= X^=^-^-^"
   ,"=#X^-J  =#"]
  -- Level 23
  ,["-J=JXJ  - "
   ,"X^X^- - = "
   ,"-^X =#X^= "
   ,"-JX^  -J-J"
   ,"- -^-JX =P"]
  -- Level 24
  ,["  -JX - X^"
   ,"XJ=^X^XJ-P"
   ,"=^XJ- - =#"
   ,"- =JX^-J=#"
   ,"=J-^X - =#"]
  -- Level 25
  ,["=^-J=#XP- "
   ,"=^- -J- =#"
   ,"=^- - =J=J"
   ,"-J- - = =J"
   ,"= -^- =J-^"]
  -- Level 26
  ,["  =^=J-PX^"
   ,"- - -^- - "
   ,"X XJ- -   "
   ,"-J  -^=#= "
   ,"  - X -J=#"]
  -- Level 27
  ,["X^XJXJ=J  "
   ,"XP- =#-^=J"
   ,"=^-^- X -^"
   ,"- - X = XJ"
   ,"XJX^  XJ= "]
  -- Level 28
  ,["-P-^  XJ=^"
   ,"-^=#  =J= "
   ,"-J-^=JXJ-J"
   ,"X -^- X^X "
   ,"= - = =#-J"]
  -- Level 29
  ,["-P=^X   -^"
   ,"X^X   - X "
   ,"=^= -J-^= "
   ,"-J-JX =#XJ"
   ,"=^X^X^- XJ"]
  -- Level 30
  ,["X^- =J=#X "
   ,"XJ=J=#X X^"
   ,"X^-J=#-^=J"
   ,"  -   - =P"
   ,"=J=#=#= - "]
  -- Level 31
  ,["- X^X -JX^"
   ,"-^XJ=#=J  "
   ,"X - X^XJ-P"
   ,"-J  =#X^XJ"
   ,"  -   =^X "]
  -- Level 32
  ,["=J- - X   "
   ,"= =^  -P=^"
   ,"=J-^=J  - "
   ,"=^=J=#- = "
   ,"-   =#=JXJ"]
  -- Level 33
  ,["= -J-J= XJ"
   ,"-J=^=#X^- "
   ,"- -J-^=J-P"
   ,"X =^=^X^XJ"
   ,"  -^-   -^"]
  -- Level 34
  ,["=J=#-JXJ-P"
   ,"XJ-^=^-J=#"
   ,"  - X - X^"
   ,"X^-^-^-^-J"
   ,"= = = =   "]
  -- Level 35
  ,["= X -^=J= "
   ,"-J-J- =^XP"
   ,"=#-^XJ-^=^"
   ,"X X^    =#"
   ,"X^X^-^X =J"]
  -- Level 36
  ,["-PXJ- -^-^"
   ,"=J=^-J-J=^"
   ,"- X^-JX^X "
   ,"-JX^-^= =#"
   ,"- X^=J  X "]
  -- Level 37
  ,["-JX =#  X^"
   ,"X^=J-J  -^"
   ,"  -^-J= = "
   ,"= =J  -P=^"
   ,"=J  X =J= "]
  -- Level 38
  ,["X^XJ=^  XJ"
   ,"-^X =#-^XP"
   ,"=J- X^- -J"
   ,"X^=JXJ- =#"
   ,"XJ= =#X =#"]
  -- Level 39
  ,["-JX   =^= "
   ,"X^-J=#-^=J"
   ,"-J-^=^XJ- "
   ,"- X XJ-PX "
   ,"=#-^=#-J=#"]
  -- Level 40
  ,["X -J= =P  "
   ,"X   =#X -J"
   ,"- -J=^X -^"
   ,"XJ-   =#  "
   ,"X -J-^- =J"]
  -- Level 41
  ,["=JXJ- - X^"
   ,"-JX   X X "
   ,"XJ=^  - =^"
   ,"X^=^= -J-P"
   ,"-   XJ=^=J"]
  -- Level 42
  ,["=JX^-J-^=P"
   ,"X XJ-^  XJ"
   ,"X = X -^-^"
   ,"-J-^X     "
   ,"X -^- =^X^"]
  -- Level 43
  ,["=J-J- -^- "
   ,"XJ=J-^  =J"
   ,"X^X = -^=^"
   ,"X^=^=^    "
   ,"-J- = -P= "]
  -- Level 44
  ,["XJ  X^-^=J"
   ,"  XJXJXJ= "
   ,"X^=^X -^-^"
   ,"-^X -J=J= "
   ,"=P=^= = =^"]
  -- Level 45
  ,["-^=^  X^-P"
   ,"XJX =   -J"
   ,"=^-^=#X^-^"
   ,"  =     XJ"
   ,"  XJXJX^=^"]
  -- Level 46
  ,["-J  -^-JX "
   ,"=^  -J=^-^"
   ,"XJ-^X^  XJ"
   ,"-^= -^X   "
   ,"-P- = = X "]
  -- Level 47
  ,["X X =P=^= "
   ,"XJ=^=#- =J"
   ,"=^-J=J=^XJ"
   ,"- =^=#X^-J"
   ,"-^=^=^=#XJ"]
  -- Level 48
  ,["= = X -J=J"
   ,"=P-^=#X X^"
   ,"=J  -^X =J"
   ,"  =^-^=   "
   ,"-^XJ= X^- "]
  -- Level 49
  ,["-J=^  X X "
   ,"-   -PXJ-^"
   ,"  X -J-^X "
   ,"-^-^= X^X "
   ,"X^X^=J=J- "]
  -- Level 50
  ,["-^X =^  -J"
   ,"=^X - =#-^"
   ,"-J-^=J= - "
   ,"X -^-^= -^"
   ,"=#-P=JXJ- "]
  -- Level 51
  ,["-P-^-JXJ-^"
   ,"XJ= =   =#"
   ,"-^=^-^X   "
   ,"-J- -J  =J"
   ,"= - = XJ-^"]
  -- Level 52
  ,["=#XJX^=JX "
   ,"= =P-J=^- "
   ,"= - =#=#X "
   ,"- X^=J- -^"
   ,"-J= =#-JX "]
  -- Level 53
  ,["XJ- -   - "
   ,"- - =^-J=#"
   ,"=^XJ-P-^-J"
   ,"-J-^-^X^X "
   ,"X = - XJ=J"]
  -- Level 54
  ,["XJ=JX^-J  "
   ,"  -JX -^= "
   ,"-J=^  - X^"
   ,"= =^= X - "
   ,"X^- -PX^X "]
  -- Level 55
  ,["X - =J- - "
   ,"- =^-P  -J"
   ,"X^-^X   =^"
   ,"X^X = X =^"
   ,"=^XJ  X XJ"]
  -- Level 56
  ,["= =^= XJ=#"
   ,"-     X XP"
   ,"- =#=J= -^"
   ,"X XJ- -J=#"
   ,"=J-^X^XJ- "]
  -- Level 57
  ,["X XJ= XJ- "
   ,"-^=#  X =#"
   ,"-J-J=#  X^"
   ,"X =^-P=^- "
   ,"=J=J- - XJ"]
  -- Level 58
  ,["X X =^=#-^"
   ,"-J-J-^-^= "
   ,"XJ-^X^=J=#"
   ,"=J=J=^- =^"
   ,"XP- X - -J"]
  -- Level 59
  ,["  X =J=^=J"
   ,"X^=#= =^  "
   ,"-^XJ-J=#=^"
   ,"X^- -^- XP"
   ,"  X^=^X^-J"]
  -- Level 60
  ,["=J= =^=#=J"
   ,"XJX = -^X "
   ,"X^=J=J-^- "
   ,"- X^=#-^=J"
   ,"- -J-P= = "]
  -- Level 61
  ,["=J  - - X "
   ,"-J=#-J=^-J"
   ,"= =#= =^= "
   ,"  XJX X^XP"
   ,"XJ=J=^=^-^"]
  -- Level 62
  ,["=^X - =^-J"
   ,"-^-^-JXJ-J"
   ,"X X -^X^-P"
   ,"X = XJ-^X "
   ,"X XJ=#X X^"]
  -- Level 63
  ,["XJ  -JX -J"
   ,"  = = =^=^"
   ,"  XJ=#-^X "
   ,"=P-^- X X^"
   ,"X^X XJ-JXJ"]
  -- Level 64
  ,["  XJXJXJ  "
   ,"XJ  XJ  XJ"
   ,"XJ  XJ  XJ"
   ,"XJ  XJXJXJ"
   ,"=PXJXJXJ  "]]

main :: IO ()
main = sequence_ $ fmap doLevel [1..(length levels)]
  where
    doLevel i = do
      putStrLn $ "Level " ++ show i
      print $ solveLevel $ levels !! (i-1)
