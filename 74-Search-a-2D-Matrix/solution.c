bool isSmallerThanThisLine(int** matrix, int matrixRowSize, int matrixColSize, int target, int lineNum) {
        return (lineNum >= 0 && lineNum < matrixRowSize) && (target <= matrix[lineNum][matrixColSize - 1]);
}
bool isInThisLine(int * line, int lineLength, int target) {
    int low = 0;
    int high = lineLength - 1;
    int mid = 0;
    while(low <= high) {
        mid = low + (high - low) / 2;
        if (line[mid] > target)
            high = mid - 1;
        else if (line[mid] < target)
            low = mid + 1;
        else
            return true;
    }
    return false;
}

bool searchMatrix(int** matrix, int matrixRowSize, int matrixColSize, int target) {
    int low = 0;
    int high = matrixRowSize;
    int mid = 0;
    while (low + 1 < high) {
        mid = low + (high - low) / 2;
        if (isSmallerThanThisLine(matrix, matrixRowSize, matrixColSize, target, mid))
            high = mid;
        else
            low = mid;
    }
    return isSmallerThanThisLine(matrix, matrixRowSize, matrixColSize, target, low) ? 
           isInThisLine(matrix[low], matrixColSize, target) : 
           (high >= matrixRowSize) ? false : isInThisLine(matrix[high], matrixColSize, target);
}